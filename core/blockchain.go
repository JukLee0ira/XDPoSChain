// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package core implements the Ethereum consensus protocol.
package core

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/XinFinOrg/XDPoSChain/XDCx/tradingstate"
	"github.com/XinFinOrg/XDPoSChain/XDCxlending/lendingstate"
	"github.com/XinFinOrg/XDPoSChain/accounts/abi/bind"
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/common/lru"
	"github.com/XinFinOrg/XDPoSChain/common/mclock"
	"github.com/XinFinOrg/XDPoSChain/common/prque"
	"github.com/XinFinOrg/XDPoSChain/common/sort"
	"github.com/XinFinOrg/XDPoSChain/consensus"
	"github.com/XinFinOrg/XDPoSChain/consensus/XDPoS"
	"github.com/XinFinOrg/XDPoSChain/consensus/XDPoS/utils"
	contractValidator "github.com/XinFinOrg/XDPoSChain/contracts/validator/contract"
	"github.com/XinFinOrg/XDPoSChain/core/rawdb"
	"github.com/XinFinOrg/XDPoSChain/core/state"
	"github.com/XinFinOrg/XDPoSChain/core/types"
	"github.com/XinFinOrg/XDPoSChain/core/vm"
	"github.com/XinFinOrg/XDPoSChain/ethclient"
	"github.com/XinFinOrg/XDPoSChain/ethdb"
	"github.com/XinFinOrg/XDPoSChain/event"
	"github.com/XinFinOrg/XDPoSChain/internal/syncx"
	"github.com/XinFinOrg/XDPoSChain/log"
	"github.com/XinFinOrg/XDPoSChain/metrics"
	"github.com/XinFinOrg/XDPoSChain/params"
	"github.com/XinFinOrg/XDPoSChain/rlp"
)

var (
	headBlockGauge     = metrics.NewRegisteredGauge("chain/head/block", nil)
	headHeaderGauge    = metrics.NewRegisteredGauge("chain/head/header", nil)
	headFastBlockGauge = metrics.NewRegisteredGauge("chain/head/receipt", nil)

	chainInfoGauge = metrics.NewRegisteredGaugeInfo("chain/info", nil)

	accountReadTimer   = metrics.NewRegisteredResettingTimer("chain/account/reads", nil)
	accountHashTimer   = metrics.NewRegisteredResettingTimer("chain/account/hashes", nil)
	accountUpdateTimer = metrics.NewRegisteredResettingTimer("chain/account/updates", nil)
	accountCommitTimer = metrics.NewRegisteredResettingTimer("chain/account/commits", nil)

	storageReadTimer   = metrics.NewRegisteredResettingTimer("chain/storage/reads", nil)
	storageHashTimer   = metrics.NewRegisteredResettingTimer("chain/storage/hashes", nil)
	storageUpdateTimer = metrics.NewRegisteredResettingTimer("chain/storage/updates", nil)
	storageCommitTimer = metrics.NewRegisteredResettingTimer("chain/storage/commits", nil)

	blockInsertTimer     = metrics.NewRegisteredResettingTimer("chain/inserts", nil)
	blockValidationTimer = metrics.NewRegisteredResettingTimer("chain/validation", nil)
	blockExecutionTimer  = metrics.NewRegisteredResettingTimer("chain/execution", nil)
	blockWriteTimer      = metrics.NewRegisteredResettingTimer("chain/write", nil)

	blockReorgMeter     = metrics.NewRegisteredMeter("chain/reorg/executes", nil)
	blockReorgAddMeter  = metrics.NewRegisteredMeter("chain/reorg/add", nil)
	blockReorgDropMeter = metrics.NewRegisteredMeter("chain/reorg/drop", nil)

	errInsertionInterrupted = errors.New("insertion is interrupted")
	errChainStopped         = errors.New("blockchain is stopped")
	errInvalidOldChain      = errors.New("invalid old chain")
	errInvalidNewChain      = errors.New("invalid new chain")

	CheckpointCh = make(chan int)
)

const (
	bodyCacheLimit      = 256
	blockCacheLimit     = 256
	receiptsCacheLimit  = 32
	maxFutureBlocks     = 256
	maxTimeFutureBlocks = 30
	badBlockLimit       = 10
	triesInMemory       = 128

	// BlockChainVersion ensures that an incompatible database forces a resync from scratch.
	//
	// Changelog:
	//
	// - Version 4
	//   The following incompatible database changes were added:
	//   * the `BlockNumber`, `TxHash`, `TxIndex`, `BlockHash` and `Index` fields of log are deleted
	//   * the `Bloom` field of receipt is deleted
	//   * the `BlockIndex` and `TxIndex` fields of txlookup are deleted
	// - Version 5
	//  The following incompatible database changes were added:
	//    * the `TxHash`, `GasCost`, and `ContractAddress` fields are no longer stored for a receipt
	//    * the `TxHash`, `GasCost`, and `ContractAddress` fields are computed by looking up the
	//      receipts' corresponding block
	// - Version 6
	//  The following incompatible database changes were added:
	//    * Transaction lookup information stores the corresponding block number instead of block hash
	BlockChainVersion uint64 = 6

	// Maximum length of chain to cache by block's number
	blocksHashCacheLimit = 900
)

// CacheConfig contains the configuration values for the trie caching/pruning
// that's resident in a blockchain.
type CacheConfig struct {
	Disabled      bool          // Whether to disable trie write caching (archive node)
	TrieNodeLimit int           // Memory limit (MB) at which to flush the current in-memory trie to disk
	TrieTimeLimit time.Duration // Time limit after which to flush the current in-memory trie to disk
}
type ResultProcessBlock struct {
	logs         []*types.Log
	receipts     []*types.Receipt
	state        *state.StateDB
	tradingState *tradingstate.TradingStateDB
	lendingState *lendingstate.LendingStateDB
	proctime     time.Duration
	usedGas      uint64
}

// BlockChain represents the canonical chain given a database with a genesis
// block. The Blockchain manages chain imports, reverts, chain reorganisations.
//
// Importing blocks in to the block chain happens according to the set of rules
// defined by the two stage Validator. Processing of blocks is done using the
// Processor which processes the included transaction. The validation of the state
// is done in the second part of the Validator. Failing results in aborting of
// the import.
//
// The BlockChain also helps in returning blocks from **any** chain included
// in the database as well as blocks that represents the canonical chain. It's
// important to note that GetBlock can return any block and does not need to be
// included in the canonical one where as GetBlockByNumber always represents the
// canonical chain.
type BlockChain struct {
	chainDb      ethdb.Database
	processor    types.BlockProcessor
	eventMux     *event.TypeMux
	genesisBlock *types.Block
	// Last known total difficulty
	mu      sync.RWMutex
	chainmu sync.RWMutex
	tsmu    sync.RWMutex

	checkpoint    int           // checkpoint counts towards the new checkpoint
	currentHeader *types.Header // Current head of the header chain (may be above the block chain!)
	currentBlock  *types.Block  // Current head of the block chain

	headerCache  *lru.Cache // Cache for the most recent block headers
	bodyCache    *lru.Cache // Cache for the most recent block bodies
	bodyRLPCache *lru.Cache // Cache for the most recent block bodies in RLP encoded format
	tdCache      *lru.Cache // Cache for the most recent block total difficulties
	blockCache   *lru.Cache // Cache for the most recent entire blocks
	futureBlocks *lru.Cache // future blocks are blocks added for later processing

	quit    chan struct{}
	running int32 // running must be called automically
	// procInterrupt must be atomically called
	procInterrupt int32 // interrupt signaler for block processing
	wg            sync.WaitGroup
	quit          chan struct{} // shutdown signal, closed in Stop.
	running       int32         // 0 if chain is running, 1 when stopped
	procInterrupt int32         // interrupt signaler for block processing

	engine    consensus.Engine
	processor Processor // block processor interface
	validator Validator // block and state validator interface
	vmConfig  vm.Config

	IPCEndpoint string
	Client      bind.ContractBackend // Global ipc client instance.

	// Blocks hash array by block number
	// cache field for tracking finality purpose, can't use for tracking block vs block relationship
	blocksHashCache *lru.Cache[uint64, []common.Hash]

	resultTrade         *lru.Cache[common.Hash, interface{}] // trades result: key - takerOrderHash, value: trades corresponding to takerOrder
	rejectedOrders      *lru.Cache[common.Hash, interface{}] // rejected orders: key - takerOrderHash, value: rejected orders corresponding to takerOrder
	resultLendingTrade  *lru.Cache[common.Hash, interface{}]
	rejectedLendingItem *lru.Cache[common.Hash, interface{}]
	finalizedTrade      *lru.Cache[common.Hash, interface{}] // include both trades which force update to closed/liquidated by the protocol
}

// NewBlockChain returns a fully initialised block chain using information
// available in the database. It initialises the default Ethereum Validator and
// Processor.
func NewBlockChain(db ethdb.Database, cacheConfig *CacheConfig, chainConfig *params.ChainConfig, engine consensus.Engine, vmConfig vm.Config) (*BlockChain, error) {
	if cacheConfig == nil {
		cacheConfig = &CacheConfig{
			TrieNodeLimit: 256 * 1024 * 1024,
			TrieTimeLimit: 5 * time.Minute,
		}
	}

	bc := &BlockChain{
		chainConfig:         chainConfig,
		cacheConfig:         cacheConfig,
		db:                  db,
		triegc:              prque.New[int64, common.Hash](nil),
		stateCache:          state.NewDatabase(db),
		quit:                make(chan struct{}),
		chainmu:             syncx.NewClosableMutex(),
		bodyCache:           lru.NewCache[common.Hash, *types.Body](bodyCacheLimit),
		bodyRLPCache:        lru.NewCache[common.Hash, rlp.RawValue](bodyCacheLimit),
		receiptsCache:       lru.NewCache[common.Hash, types.Receipts](receiptsCacheLimit),
		blockCache:          lru.NewCache[common.Hash, *types.Block](blockCacheLimit),
		futureBlocks:        lru.NewCache[common.Hash, *types.Block](maxFutureBlocks),
		resultProcess:       lru.NewCache[common.Hash, *ResultProcessBlock](blockCacheLimit),
		calculatingBlock:    lru.NewCache[common.Hash, *CalculatedBlock](blockCacheLimit),
		downloadingBlock:    lru.NewCache[common.Hash, struct{}](blockCacheLimit),
		engine:              engine,
		vmConfig:            vmConfig,
		badBlocks:           lru.NewCache[common.Hash, *types.Header](badBlockLimit),
		blocksHashCache:     lru.NewCache[uint64, []common.Hash](blocksHashCacheLimit),
		resultTrade:         lru.NewCache[common.Hash, interface{}](tradingstate.OrderCacheLimit),
		rejectedOrders:      lru.NewCache[common.Hash, interface{}](tradingstate.OrderCacheLimit),
		resultLendingTrade:  lru.NewCache[common.Hash, interface{}](tradingstate.OrderCacheLimit),
		rejectedLendingItem: lru.NewCache[common.Hash, interface{}](tradingstate.OrderCacheLimit),
		finalizedTrade:      lru.NewCache[common.Hash, interface{}](tradingstate.OrderCacheLimit),
	}
	bc.SetValidator(NewBlockValidator(chainConfig, bc, engine))
	bc.SetProcessor(NewStateProcessor(chainConfig, bc, engine))

	var err error
	bc.hc, err = NewHeaderChain(db, chainConfig, engine, bc.getProcInterrupt)
	if err != nil {
		return nil, err
	}
	bc.genesisBlock = bc.GetBlockByNumber(0)
	if bc.genesisBlock == nil {
		reader, err := NewDefaultGenesisReader()
		if err != nil {
			return nil, err
		}
		bc.genesisBlock, err = WriteGenesisBlock(chainDb, reader)
		if err != nil {
			return nil, err
		}
		glog.V(logger.Info).Infoln("WARNING: Wrote default ethereum genesis block")
	}
	if err := bc.loadLastState(); err != nil {
		return nil, err
	}
	// Check the current state of the block hashes and make sure that we do not have any of the bad blocks in our chain
	for hash, _ := range BadHashes {
		if header := bc.GetHeader(hash); header != nil {
			glog.V(logger.Error).Infof("Found bad hash, rewinding chain to block #%d [%x…]", header.Number, header.ParentHash[:4])
			bc.SetHead(header.Number.Uint64() - 1)
			glog.V(logger.Error).Infoln("Chain rewind was successful, resuming normal operation")
		}
	}
	// Take ownership of this particular state
	go bc.update()
	return bc, nil
}

// loadLastState loads the last known chain state from the database. This method
// assumes that the chain manager mutex is held.
func (self *BlockChain) loadLastState() error {
	// Restore the last known head block
	head := GetHeadBlockHash(self.chainDb)
	if head == (common.Hash{}) {
		// Corrupt or empty database, init from scratch
		self.Reset()
	} else {
		if block := self.GetBlock(head); block != nil {
			// Block found, set as the current head
			self.currentBlock = block
		} else {
			// Corrupt or empty database, init from scratch
			self.Reset()
		}
	}
	// Restore the last known head header
	self.currentHeader = self.currentBlock.Header()
	if head := GetHeadHeaderHash(self.chainDb); head != (common.Hash{}) {
		if header := self.GetHeader(head); header != nil {
			self.currentHeader = header
		}
	}
	// Issue a status log and return
	headerTd := self.GetTd(self.currentHeader.Hash())
	blockTd := self.GetTd(self.currentBlock.Hash())

	glog.V(logger.Info).Infof("Last header: #%d [%x…] TD=%v", self.currentHeader.Number, self.currentHeader.Hash(), headerTd)
	glog.V(logger.Info).Infof("Last block: #%d [%x…] TD=%v", self.currentBlock.Number(), self.currentBlock.Hash(), blockTd)

	return nil
}

// SetHead rewind the local chain to a new head entity. In the case of headers,
// everything above the new head will be deleted and the new one set. In the case
// of blocks though, the head may be further rewound if block bodies are missing
// (non-archive nodes after a fast sync).
func (bc *BlockChain) SetHead(head uint64) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// Delete everything from the current header head (is above block head)
	for i := bc.currentHeader.Number.Uint64(); i > head; i-- {
		if hash := GetCanonicalHash(bc.chainDb, i); hash != (common.Hash{}) {
			DeleteCanonicalHash(bc.chainDb, i)
			DeleteHeader(bc.chainDb, hash)
			DeleteBody(bc.chainDb, hash)
			DeleteTd(bc.chainDb, hash)
		}
	}
	bc.currentHeader = GetHeader(bc.chainDb, GetCanonicalHash(bc.chainDb, head))

	// Rewind the block chain until a whole block is found
	for bc.GetBlockByNumber(head) == nil {
		head--
	}
	bc.currentBlock = bc.GetBlockByNumber(head)

	// Clear out any stale content from the caches
	bc.headerCache.Purge()
	bc.bodyCache.Purge()
	bc.bodyRLPCache.Purge()
	bc.receiptsCache.Purge()
	bc.blockCache.Purge()
	bc.futureBlocks.Purge()

	// Update all computed fields to the new head
	bc.insert(bc.currentBlock)
	bc.loadLastState()
}

func (self *BlockChain) GasLimit() *big.Int {
	self.mu.RLock()
	defer self.mu.RUnlock()

	return self.currentBlock.GasLimit()
}

func (self *BlockChain) LastBlockHash() common.Hash {
	self.mu.RLock()
	defer self.mu.RUnlock()

	return self.currentBlock.Hash()
}

// CurrentHeader retrieves the current head header of the canonical chain. The
// header is retrieved from the chain manager's internal cache, involving no
// database operations.
func (self *BlockChain) CurrentHeader() *types.Header {
	self.mu.RLock()
	defer self.mu.RUnlock()

	return self.currentHeader
}

// CurrentBlock retrieves the current head block of the canonical chain. The
// block is retrieved from the chain manager's internal cache, involving no
// database operations.
func (self *BlockChain) CurrentBlock() *types.Block {
	self.mu.RLock()
	defer self.mu.RUnlock()

	return self.currentBlock
}

func (self *BlockChain) Status() (td *big.Int, currentBlock common.Hash, genesisBlock common.Hash) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	return self.GetTd(self.currentBlock.Hash()), self.currentBlock.Hash(), self.genesisBlock.Hash()
}

func (self *BlockChain) SetProcessor(proc types.BlockProcessor) {
	self.processor = proc
}

func (self *BlockChain) State() (*state.StateDB, error) {
	return state.New(self.CurrentBlock().Root(), self.chainDb)
}

// Reset purges the entire blockchain, restoring it to its genesis state.
func (bc *BlockChain) Reset() error {
	return bc.ResetWithGenesisBlock(bc.genesisBlock)
}

// ResetWithGenesisBlock purges the entire blockchain, restoring it to the
// specified genesis state.
func (bc *BlockChain) ResetWithGenesisBlock(genesis *types.Block) error {
	// Dump the entire block chain and purge the caches
	for block := bc.currentBlock; block != nil; block = bc.GetBlock(block.ParentHash()) {
		DeleteBlock(bc.chainDb, block.Hash())
	}
	for header := bc.currentHeader; header != nil; header = bc.GetHeader(header.ParentHash) {
		DeleteBlock(bc.chainDb, header.Hash())
	}
	bc.headerCache.Purge()
	bc.bodyCache.Purge()
	bc.bodyRLPCache.Purge()
	bc.blockCache.Purge()
	bc.futureBlocks.Purge()

	// Prepare the genesis block and reinitialise the chain
	batch := bc.db.NewBatch()
	rawdb.WriteTd(batch, genesis.Hash(), genesis.NumberU64(), genesis.Difficulty())
	rawdb.WriteBlock(batch, genesis)
	if err := batch.Write(); err != nil {
		log.Crit("Failed to write genesis block", "err", err)
	}
	bc.writeHeadBlock(genesis, false)

	// Last update all in-memory chain markers
	bc.genesisBlock = genesis
	bc.insert(bc.genesisBlock)
	bc.currentBlock = bc.genesisBlock
	bc.currentHeader = bc.genesisBlock.Header()
}

// repair tries to repair the current blockchain by rolling back the current block
// until one with associated state is found. This is needed to fix incomplete db
// writes caused either by crashes/power outages, or simply non-committed tries.
//
// This method only rolls back the current block. The current header and current
// fast block are left intact.
func (bc *BlockChain) repair(head **types.Block) error {
	for {
		// Abort if we've rewound to a head block that does have associated state
		if (common.Rewound == uint64(0)) || ((*head).Number().Uint64() < common.Rewound) {
			if _, err := state.New((*head).Root(), bc.stateCache); err == nil {
				log.Info("Rewound blockchain to past state", "number", (*head).Number(), "hash", (*head).Hash())
				engine, ok := bc.Engine().(*XDPoS.XDPoS)
				if ok {
					tradingService := engine.GetXDCXService()
					lendingService := engine.GetLendingService()
					if bc.Config().IsTIPXDCXReceiver((*head).Number()) && bc.chainConfig.XDPoS != nil && (*head).NumberU64() > bc.chainConfig.XDPoS.Epoch && tradingService != nil && lendingService != nil {
						author, _ := bc.Engine().Author((*head).Header())
						tradingRoot, err := tradingService.GetTradingStateRoot(*head, author)
						if err == nil {
							_, err = tradingstate.New(tradingRoot, tradingService.GetStateCache())
						}
						if err == nil {
							lendingRoot, err := lendingService.GetLendingStateRoot(*head, author)
							if err == nil {
								_, err = lendingstate.New(lendingRoot, lendingService.GetStateCache())
								if err == nil {
									return nil
								}
							}
						}
					} else {
						return nil
					}
				} else {
					return nil
				}
			}
		} else {
			log.Info("Rewound blockchain to past state", "number", (*head).Number(), "hash", (*head).Hash())
		}
		// Otherwise rewind one block and recheck state availability there
		(*head) = bc.GetBlock((*head).ParentHash(), (*head).NumberU64()-1)
	}
}

// Export writes the active chain to the given writer.
func (bc *BlockChain) Export(w io.Writer) error {
	return bc.ExportN(w, uint64(0), bc.CurrentBlock().NumberU64())
}

// ExportN writes a subset of the active chain to the given writer.
func (bc *BlockChain) ExportN(w io.Writer, first uint64, last uint64) error {
	if !bc.chainmu.TryLock() {
		return errChainStopped
	}
	defer bc.chainmu.Unlock()

	if first > last {
		return fmt.Errorf("export failed: first (%d) is greater than last (%d)", first, last)
	}
	log.Info("Exporting batch of blocks", "count", last-first+1)

	for nr := first; nr <= last; nr++ {
		block := bc.GetBlockByNumber(nr)
		if block == nil {
			return fmt.Errorf("export failed on #%d: not found", nr)
		}

		if err := block.EncodeRLP(w); err != nil {
			return err
		}
	}
	return nil
}

// insert injects a new head block into the current block chain. This method
// assumes that the block is indeed a true head. It will also reset the head
// header to this very same block to prevent the headers from diverging on a
// different header chain.
//
// Note, this function assumes that the `mu` mutex is held!
func (bc *BlockChain) insert(block *types.Block) {
	// Add the block to the canonical chain number scheme and mark as the head
	if err := WriteCanonicalHash(bc.chainDb, block.Hash(), block.NumberU64()); err != nil {
		glog.Fatalf("failed to insert block number: %v", err)
	}
	if err := WriteHeadBlockHash(bc.chainDb, block.Hash()); err != nil {
		glog.Fatalf("failed to insert head block hash: %v", err)
	}
	if err := WriteHeadHeaderHash(bc.chainDb, block.Hash()); err != nil {
		glog.Fatalf("failed to insert head header hash: %v", err)
	}
	// Update the internal state with the head block
	bc.currentBlock = block
	bc.currentHeader = block.Header()
}

// Accessors
func (bc *BlockChain) Genesis() *types.Block {
	return bc.genesisBlock
}

// HasHeader checks if a block header is present in the database or not, caching
// it if present.
func (bc *BlockChain) HasHeader(hash common.Hash) bool {
	return bc.GetHeader(hash) != nil
}

// GetHeader retrieves a block header from the database by hash, caching it if
// found.
func (self *BlockChain) GetHeader(hash common.Hash) *types.Header {
	// Short circuit if the header's already in the cache, retrieve otherwise
	if header, ok := self.headerCache.Get(hash); ok {
		return header.(*types.Header)
	}
	header := GetHeader(self.chainDb, hash)
	if header == nil {
		return nil
	}
	// Cache the found header for next time and return
	self.headerCache.Add(header.Hash(), header)
	return header
}

// Genesis retrieves the chain's genesis block.
func (bc *BlockChain) Genesis() *types.Block {
	return bc.genesisBlock
}

// GetBody retrieves a block body (transactions and uncles) from the database by
// hash, caching it if found.
func (bc *BlockChain) GetBody(hash common.Hash) *types.Body {
	// Short circuit if the body's already in the cache, retrieve otherwise
	if cached, ok := bc.bodyCache.Get(hash); ok {
		return cached
	}
	number := bc.hc.GetBlockNumber(hash)
	if number == nil {
		return nil
	}
	body := rawdb.ReadBody(bc.db, hash, *number)
	if body == nil {
		return nil
	}
	// Cache the found body for next time and return
	bc.bodyCache.Add(hash, body)
	return body
}

// GetBodyRLP retrieves a block body in RLP encoding from the database by hash,
// caching it if found.
func (bc *BlockChain) GetBodyRLP(hash common.Hash) rlp.RawValue {
	// Short circuit if the body's already in the cache, retrieve otherwise
	if cached, ok := bc.bodyRLPCache.Get(hash); ok {
		return cached
	}
	number := bc.hc.GetBlockNumber(hash)
	if number == nil {
		return nil
	}
	body := rawdb.ReadBodyRLP(bc.db, hash, *number)
	if len(body) == 0 {
		return nil
	}
	// Cache the found body for next time and return
	bc.bodyRLPCache.Add(hash, body)
	return body
}

// HasBlock checks if a block is fully present in the database or not.
func (bc *BlockChain) HasBlock(hash common.Hash, number uint64) bool {
	if bc.blockCache.Contains(hash) {
		return true
	}
	if !bc.HasHeader(hash, number) {
		return false
	}
	return rawdb.HasBody(bc.db, hash, number)
}

// HasFullState checks if state trie is fully present in the database or not.
func (bc *BlockChain) HasFullState(block *types.Block) bool {
	_, err := bc.stateCache.OpenTrie(block.Root())
	if err != nil {
		return false
	}
	engine, _ := bc.Engine().(*XDPoS.XDPoS)
	if bc.Config().IsTIPXDCX(block.Number()) && bc.chainConfig.XDPoS != nil && engine != nil && block.NumberU64() > bc.chainConfig.XDPoS.Epoch {
		tradingService := engine.GetXDCXService()
		lendingService := engine.GetLendingService()
		author, _ := bc.Engine().Author(block.Header())
		if tradingService != nil && !tradingService.HasTradingState(block, author) {
			return false
		}
		if lendingService != nil && !lendingService.HasLendingState(block, author) {
			return false
		}
	}
	return true
}

// HasBlockAndFullState checks if a block and associated state trie is fully present
// in the database or not, caching it if present.
func (bc *BlockChain) HasBlockAndFullState(hash common.Hash, number uint64) bool {
	// Check first that the block itself is known
	block := bc.GetBlock(hash, number)
	if block == nil {
		return false
	}
	return bc.HasFullState(block)
}

// GetBlock retrieves a block from the database by hash and number,
// caching it if found.
func (bc *BlockChain) GetBlock(hash common.Hash, number uint64) *types.Block {
	// Short circuit if the block's already in the cache, retrieve otherwise
	if block, ok := bc.blockCache.Get(hash); ok {
		return block
	}
	block := rawdb.ReadBlock(bc.db, hash, number)
	if block == nil {
		return nil
	}
	// Cache the found block for next time and return
	bc.blockCache.Add(block.Hash(), block)
	return block
}

// GetBlockByHash retrieves a block from the database by hash, caching it if found.
func (bc *BlockChain) GetBlockByHash(hash common.Hash) *types.Block {
	number := bc.hc.GetBlockNumber(hash)
	if number == nil {
		return nil
	}
	return bc.GetBlock(hash, *number)
}

// GetBlockByNumber retrieves a block from the database by number, caching it
// (associated with its hash) if found.
func (bc *BlockChain) GetBlockByNumber(number uint64) *types.Block {
	hash := rawdb.ReadCanonicalHash(bc.db, number)
	if hash == (common.Hash{}) {
		return nil
	}
	return bc.GetBlock(hash, number)
}

// GetReceiptsByHash retrieves the receipts for all transactions in a given block.
func (bc *BlockChain) GetReceiptsByHash(hash common.Hash) types.Receipts {
	if receipts, ok := bc.receiptsCache.Get(hash); ok {
		return receipts
	}
	number := rawdb.ReadHeaderNumber(bc.db, hash)
	if number == nil {
		return nil
	}
	receipts := rawdb.ReadReceipts(bc.db, hash, *number, bc.chainConfig)
	if receipts == nil {
		return nil
	}
	bc.receiptsCache.Add(hash, receipts)
	return receipts
}

// GetBlocksFromHash returns the block corresponding to hash and up to n-1 ancestors.
// [deprecated by eth/62]
func (bc *BlockChain) GetBlocksFromHash(hash common.Hash, n int) (blocks []*types.Block) {
	number := bc.hc.GetBlockNumber(hash)
	if number == nil {
		return nil
	}
	for i := 0; i < n; i++ {
		block := bc.GetBlock(hash, *number)
		if block == nil {
			break
		}
		blocks = append(blocks, block)
		hash = block.ParentHash()
		*number--
	}
	return
}

// GetUnclesInChain retrieves all the uncles from a given block backwards until
// a specific distance is reached.
func (self *BlockChain) GetUnclesInChain(block *types.Block, length int) []*types.Header {
	uncles := []*types.Header{}
	for i := 0; block != nil && i < length; i++ {
		uncles = append(uncles, block.Uncles()...)
		block = bc.GetBlock(block.ParentHash(), block.NumberU64()-1)
	}
	return uncles
}

func (bc *BlockChain) Stop() {
	if !atomic.CompareAndSwapInt32(&bc.running, 0, 1) {
		return
	}

	// Unsubscribe all subscriptions registered from blockchain.
	bc.scope.Close()

	// Signal shutdown to all goroutines.
	close(bc.quit)
	bc.StopInsert()

	// Now wait for all chain modifications to end and persistent goroutines to exit.
	//
	// Note: Close waits for the mutex to become available, i.e. any running chain
	// modification will have exited when Close returns. Since we also called StopInsert,
	// the mutex should become available quickly. It cannot be taken again after Close has
	// returned.
	bc.chainmu.Close()
	bc.wg.Wait()
	bc.saveData()
	log.Info("Blockchain manager stopped")
}

// StopInsert interrupts all insertion methods, causing them to return
// errInsertionInterrupted as soon as possible. Insertion is permanently disabled after
// calling this method.
func (bc *BlockChain) StopInsert() {
	atomic.StoreInt32(&bc.procInterrupt, 1)
}

// insertStopped returns true after StopInsert has been called.
func (bc *BlockChain) insertStopped() bool {
	return atomic.LoadInt32(&bc.procInterrupt) == 1
}

func (bc *BlockChain) procFutureBlocks() {
	blocks := make([]*types.Block, 0, bc.futureBlocks.Len())
	for _, hash := range bc.futureBlocks.Keys() {
		if block, exist := bc.futureBlocks.Peek(hash); exist {
			blocks = append(blocks, block)
		}
	}
	if len(blocks) > 0 {
		types.BlockBy(types.Number).Sort(blocks)

		// Insert one by one as chain insertion needs contiguous ancestry between blocks
		for i := range blocks {
			_, err := bc.InsertChain(blocks[i : i+1])
			// let consensus engine handle the last block (e.g. for voting)
			if i == len(blocks)-1 && err == nil {
				engine, ok := bc.Engine().(*XDPoS.XDPoS)
				if ok {
					j := i
					go func() {
						header := blocks[j].Header()
						err = engine.HandleProposedBlock(bc, header)
						if err != nil {
							log.Info("[procFutureBlocks] handle proposed block has error", "err", err, "block hash", header.Hash(), "number", header.Number)
						}
					}()
				}
			}
		}
	}
}

// WriteStatus status of write
type WriteStatus byte

const (
	NonStatTy WriteStatus = iota
	CanonStatTy
	SideStatTy
)

// writeHeader writes a header into the local chain, given that its parent is
// already known. If the total difficulty of the newly inserted header becomes
// greater than the old known TD, the canonical chain is re-routed.
//
// Note: This method is not concurrent-safe with inserting blocks simultaneously
// into the chain, as side effects caused by reorganizations cannot be emulated
// without the real blocks. Hence, writing headers directly should only be done
// in two scenarios: pure-header mode of operation (light clients), or properly
// separated header/block phases (non-archive clients).
func (self *BlockChain) writeHeader(header *types.Header) error {
	self.wg.Add(1)
	defer self.wg.Done()

	// Calculate the total difficulty of the header
	ptd := self.GetTd(header.ParentHash)
	if ptd == nil {
		return ParentError(header.ParentHash)
	}
	td := new(big.Int).Add(header.Difficulty, ptd)

	// Make sure no inconsistent state is leaked during insertion
	self.mu.Lock()
	defer self.mu.Unlock()

	// If the total difficulty is higher than our known, add it to the canonical chain
	if td.Cmp(self.GetTd(self.currentHeader.Hash())) > 0 {
		// Delete any canonical number assignments above the new head
		for i := header.Number.Uint64() + 1; GetCanonicalHash(self.chainDb, i) != (common.Hash{}); i++ {
			DeleteCanonicalHash(self.chainDb, i)
		}
		// Overwrite any stale canonical number assignments
		head := self.GetHeader(header.ParentHash)
		for GetCanonicalHash(self.chainDb, head.Number.Uint64()) != head.Hash() {
			WriteCanonicalHash(self.chainDb, head.Hash(), head.Number.Uint64())
			head = self.GetHeader(head.ParentHash)
		}
		// Extend the canonical chain with the new header
		if err := WriteCanonicalHash(self.chainDb, header.Hash(), header.Number.Uint64()); err != nil {
			glog.Fatalf("failed to insert header number: %v", err)
		}
		if err := WriteHeadHeaderHash(self.chainDb, header.Hash()); err != nil {
			glog.Fatalf("failed to insert head header hash: %v", err)
		}
		self.currentHeader = types.CopyHeader(header)
	}
	// Irrelevant of the canonical status, write the header itself to the database
	if err := WriteTd(self.chainDb, header.Hash(), td); err != nil {
		glog.Fatalf("failed to write header total difficulty: %v", err)
	}
	if err := WriteHeader(self.chainDb, header); err != nil {
		glog.Fatalf("filed to write header contents: %v", err)
	}
	return nil
}

// InsertHeaderChain will attempt to insert the given header chain in to the
// local chain, possibly creating a dork. If an error is returned,  it will
// return the index number of the failing header as well an error describing
// what went wrong.
//
// The verify parameter can be used to fine tune whether nonce verification
// should be done or not. The reason behind the optional check is because some
// of the header retrieval mechanisms already need to verfy nonces, as well as
// because nonces can be verified sparsely, not needing to check each.
func (self *BlockChain) InsertHeaderChain(chain []*types.Header, verify bool) (int, error) {
	self.wg.Add(1)
	defer self.wg.Done()

	// Make sure only one thread manipulates the chain at once
	self.chainmu.Lock()
	defer self.chainmu.Unlock()

	// Collect some import statistics to report on
	stats := struct{ processed, ignored int }{}
	start := time.Now()

	// Start the parallel nonce verifier, with a fake nonce if not requested
	verifier := self.pow
	if !verify {
		verifier = FakePow{}
	}
	nonceAbort, nonceResults := verifyNoncesFromHeaders(verifier, chain)
	defer close(nonceAbort)

	// Iterate over the headers, inserting any new ones
	complete := make([]bool, len(chain))
	for i, header := range chain {
		// Short circuit insertion if shutting down
		if atomic.LoadInt32(&self.procInterrupt) == 1 {
			glog.V(logger.Debug).Infoln("Premature abort during header chain processing")
			break
		}
		hash := header.Hash()

		// Accumulate verification results until the next header is verified
		for !complete[i] {
			if res := <-nonceResults; res.valid {
				complete[res.index] = true
			} else {
				header := chain[res.index]
				return res.index, &BlockNonceErr{
					Hash:   header.Hash(),
					Number: new(big.Int).Set(header.Number),
					Nonce:  header.Nonce.Uint64(),
				}
			}
		}
		if BadHashes[hash] {
			glog.V(logger.Error).Infof("Bad header %d [%x…], known bad hash", header.Number, hash)
			return i, BadHashError(hash)
		}
		// Write the header to the chain and get the status
		if self.HasHeader(hash) {
			stats.ignored++
			continue
		}
		if err := self.writeHeader(header); err != nil {
			return i, err
		}
		stats.processed++
	}
	// Report some public statistics so the user has a clue what's going on
	first, last := chain[0], chain[len(chain)-1]
	glog.V(logger.Info).Infof("imported %d header(s) (%d ignored) in %v. #%v [%x… / %x…]", stats.processed, stats.ignored,
		time.Since(start), last.Number, first.Hash().Bytes()[:4], last.Hash().Bytes()[:4])

	return 0, nil
}

// WriteBlock writes the block to the chain.
func (self *BlockChain) WriteBlock(block *types.Block) (status writeStatus, err error) {
	self.wg.Add(1)
	defer self.wg.Done()

	// Calculate the total difficulty of the block
	ptd := bc.GetTd(block.ParentHash(), block.NumberU64()-1)
	if ptd == nil {
		return NonStatTy, consensus.ErrUnknownAncestor
	}
	// Make sure no inconsistent state is leaked during insertion
	currentBlock := bc.CurrentBlock()
	localTd := bc.GetTd(currentBlock.Hash(), currentBlock.NumberU64())
	externTd := new(big.Int).Add(block.Difficulty(), ptd)

	self.mu.RLock()
	cblock := self.currentBlock
	self.mu.RUnlock()

	// Compare the TD of the last known block in the canonical chain to make sure it's greater.
	// At this point it's possible that a different chain (fork) becomes the new canonical chain.
	if td.Cmp(self.GetTd(self.currentBlock.Hash())) > 0 {
		// chain fork
		if block.ParentHash() != cblock.Hash() {
			// during split we merge two different chains and create the new canonical chain
			err := self.reorg(cblock, block)
			if err != nil {
				return NonStatTy, err
			}
		}
		status = CanonStatTy

		self.mu.Lock()
		self.insert(block)
		self.mu.Unlock()
	} else {
		status = SideStatTy
	}

	// Set new head.
	if status == CanonStatTy {
		// WriteBlock has already been called, no need to write again
		bc.writeHeadBlock(block, false)
		// prepare set of masternodes for the next epoch
		if bc.chainConfig.XDPoS != nil && ((block.NumberU64() % bc.chainConfig.XDPoS.Epoch) == (bc.chainConfig.XDPoS.Epoch - bc.chainConfig.XDPoS.Gap)) {
			if err := bc.UpdateM1(); err != nil {
				log.Crit("Fail to update masternodes during writeBlockWithState", "number", block.Number, "hash", block.Hash().Hex(), "err", err)
			}
		}
	}
	// save cache BlockSigners
	if bc.chainConfig.XDPoS != nil && bc.chainConfig.IsTIPSigning(block.Number()) {
		engine, ok := bc.Engine().(*XDPoS.XDPoS)
		if ok {
			engine.CacheSigningTxs(block.Header().Hash(), block.Transactions())
		}
	}
	bc.futureBlocks.Remove(block.Hash())
	return status, nil
}

// InsertChain attempts to insert the given batch of blocks in to the canonical
// chain or, otherwise, create a fork. If an error is returned it will return
// the index number of the failing block as well an error describing what went
// wrong.
//
// After insertion is done, all accumulated events will be fired.
func (bc *BlockChain) InsertChain(chain types.Blocks) (int, error) {
	// Sanity check that we have something meaningful to import
	if len(chain) == 0 {
		return 0, nil
	}

	// Do a sanity check that the provided chain is actually ordered and linked
	for i := 1; i < len(chain); i++ {
		block, prev := chain[i], chain[i-1]
		if block.NumberU64() != chain[i-1].NumberU64()+1 || block.ParentHash() != chain[i-1].Hash() {
			// Chain broke ancestry, log a messge (programming error) and skip insertion
			log.Error("Non contiguous block insert",
				"number", block.Number(),
				"hash", block.Hash(),
				"parent", block.ParentHash(),
				"prevnumber", prev.Number(),
				"prevhash", prev.Hash())

			return 0, fmt.Errorf("non contiguous insert: item %d is #%d [%x..], item %d is #%d [%x..] (parent [%x..])", i-1, prev.NumberU64(),
				prev.Hash().Bytes()[:4], i, block.NumberU64(), block.Hash().Bytes()[:4], block.ParentHash().Bytes()[:4])
		}
	}

	// Pre-check passed, start the full block imports.
	if !bc.chainmu.TryLock() {
		return 0, errChainStopped
	}
	defer bc.chainmu.Unlock()
	n, events, logs, err := bc.insertChain(chain, true)
	bc.PostChainEvents(events, logs)
	return n, err
}

// insertChain is the internal implementation of InsertChain, which assumes that
// 1) chains are contiguous, and 2) The chain mutex is held.
//
// This method is split out so that import batches that require re-injecting
// historical blocks can do so without releasing the lock, which could lead to
// racey behaviour. If a sidechain import is in progress, and the historic state
// is imported, but then new canon-head is added before the actual sidechain
// completes, then the historic state could be pruned again
func (bc *BlockChain) insertChain(chain types.Blocks, verifySeals bool) (int, []interface{}, []*types.Log, error) {
	// If the chain is terminating, don't even bother starting up.
	if bc.insertStopped() {
		return 0, nil, nil, nil
	}

	// A queued approach to delivering events. This is generally
	// faster than direct delivery and requires much less mutex
	// acquiring.
	var (
		stats         = insertStats{startTime: mclock.Now()}
		events        = make([]interface{}, 0, len(chain))
		lastCanon     *types.Block
		coalescedLogs []*types.Log
	)
	// Start the parallel header verifier
	headers := make([]*types.Header, len(chain))
	seals := make([]bool, len(chain))

	for i, block := range chain {
		headers[i] = block.Header()
		seals[i] = verifySeals
		bc.downloadingBlock.Add(block.Hash(), struct{}{})
	}
	abort, results := bc.engine.VerifyHeaders(bc, headers, seals)
	defer close(abort)

	// Start a parallel signature recovery (signer will fluke on fork transition, minimal perf loss)
	SenderCacher.RecoverFromBlocks(types.MakeSigner(bc.chainConfig, chain[0].Number()), chain)

	// Iterate over the blocks and insert when the verifier permits
	for i, block := range chain {
		if atomic.LoadInt32(&self.procInterrupt) == 1 {
			glog.V(logger.Debug).Infoln("Premature abort during block chain processing")
			break
		}
		// If the header is a banned one, straight out abort
		if BadHashes[block.Hash()] {
			bc.reportBlock(block, nil, ErrBlacklistedHash)
			return i, events, coalescedLogs, ErrBlacklistedHash
		}
		// Wait for the block's verification to complete
		bstart := time.Now()

		err := <-results
		if err == nil {
			err = bc.Validator().ValidateBody(block)
		}
		switch {
		case err == ErrKnownBlock:
			// Block and state both already known. However if the current block is below
			// this number we did a rollback and we should reimport it nonetheless.
			if bc.CurrentBlock().NumberU64() >= block.NumberU64() {
				stats.ignored++
				continue
			}

		case err == consensus.ErrFutureBlock:
			// Allow up to MaxFuture second in the future blocks. If this limit is exceeded
			// the chain is discarded and processed at a later time if given.
			max := big.NewInt(time.Now().Unix() + maxTimeFutureBlocks)
			if block.Time().Cmp(max) > 0 {
				return i, events, coalescedLogs, fmt.Errorf("future block: %v > %v", block.Time(), max)
			}
			bc.futureBlocks.Add(block.Hash(), block)
			stats.queued++
			continue

		case err == consensus.ErrUnknownAncestor && bc.futureBlocks.Contains(block.ParentHash()):
			bc.futureBlocks.Add(block.Hash(), block)
			stats.queued++
			continue

		case err == consensus.ErrPrunedAncestor:
			// Block competing with the canonical chain, store in the db, but don't process
			// until the competitor TD goes above the canonical TD
			currentBlock := bc.CurrentBlock()
			localTd := bc.GetTd(currentBlock.Hash(), currentBlock.NumberU64())
			externTd := new(big.Int).Add(bc.GetTd(block.ParentHash(), block.NumberU64()-1), block.Difficulty())
			if localTd.Cmp(externTd) > 0 {
				if err = bc.writeBlockWithoutState(block, externTd); err != nil {
					return i, events, coalescedLogs, err
				}
				continue
			}
			// Competitor chain beat canonical, gather all blocks from the common ancestor
			var winner []*types.Block

			parent := bc.GetBlock(block.ParentHash(), block.NumberU64()-1)
			for !bc.HasFullState(parent) {
				winner = append(winner, parent)
				parent = bc.GetBlock(parent.ParentHash(), parent.NumberU64()-1)
			}
			for j := 0; j < len(winner)/2; j++ {
				winner[j], winner[len(winner)-1-j] = winner[len(winner)-1-j], winner[j]
			}
			log.Debug("Number block need calculated again", "number", block.NumberU64(), "hash", block.Hash().Hex(), "winners", len(winner))
			// Import all the pruned blocks to make the state available
			// During reorg, we use verifySeals=false
			_, evs, logs, err := bc.insertChain(winner, false)
			events, coalescedLogs = evs, logs

			if err != nil {
				return i, events, coalescedLogs, err
			}

		case err != nil:
			bc.reportBlock(block, nil, err)
			return i, events, coalescedLogs, err
		}
		// Create a new statedb using the parent block and report an
		// error if it fails.
		var parent *types.Block
		if i == 0 {
			parent = bc.GetBlock(block.ParentHash(), block.NumberU64()-1)
		} else {
			parent = chain[i-1]
		}
		statedb, err := state.New(parent.Root(), bc.stateCache)
		if err != nil {
			return i, events, coalescedLogs, err
		}
		// clear the previous dry-run cache
		var tradingState *tradingstate.TradingStateDB
		var lendingState *lendingstate.LendingStateDB
		var tradingService utils.TradingService
		var lendingService utils.LendingService
		isSDKNode := false
		engine, _ := bc.Engine().(*XDPoS.XDPoS)
		if bc.Config().IsTIPXDCXReceiver(block.Number()) && bc.chainConfig.XDPoS != nil && engine != nil && block.NumberU64() > bc.chainConfig.XDPoS.Epoch {
			author, err := bc.Engine().Author(block.Header()) // Ignore error, we're past header validation
			if err != nil {
				bc.reportBlock(block, nil, err)
				return i, events, coalescedLogs, err
			}
			parentAuthor, _ := bc.Engine().Author(parent.Header())
			tradingService = engine.GetXDCXService()
			lendingService = engine.GetLendingService()
			if tradingService != nil && lendingService != nil {
				isSDKNode = tradingService.IsSDKNode()
				txMatchBatchData, err := ExtractTradingTransactions(block.Transactions())
				if err != nil {
					bc.reportBlock(block, nil, err)
					return i, events, coalescedLogs, err
				}
				tradingState, err = tradingService.GetTradingState(parent, parentAuthor)
				if err != nil {
					bc.reportBlock(block, nil, err)
					return i, events, coalescedLogs, err
				}
				lendingState, err = lendingService.GetLendingState(parent, parentAuthor)
				if err != nil {
					bc.reportBlock(block, nil, err)
					return i, events, coalescedLogs, err
				}
				isEpochSwithBlock, epochNumber, err := engine.IsEpochSwitch(block.Header())
				if err != nil {
					log.Error("[insertChain] Error while checking if the incoming block is epoch switch block", "Hash", block.Hash(), "Number", block.Number())
					bc.reportBlock(block, nil, err)
				}
				if isEpochSwithBlock {
					if err := tradingService.UpdateMediumPriceBeforeEpoch(epochNumber, tradingState, statedb); err != nil {
						return i, events, coalescedLogs, err
					}
				} else {
					for _, txMatchBatch := range txMatchBatchData {
						log.Debug("Verify matching transaction", "txHash", txMatchBatch.TxHash.Hex())
						err := bc.Validator().ValidateTradingOrder(statedb, tradingState, txMatchBatch, author, block.Header())
						if err != nil {
							bc.reportBlock(block, nil, err)
							return i, events, coalescedLogs, err
						}
					}
					//
					batches, err := ExtractLendingTransactions(block.Transactions())
					if err != nil {
						bc.reportBlock(block, nil, err)
						return i, events, coalescedLogs, err
					}
					for _, batch := range batches {
						log.Debug("Verify matching transaction", "txHash", batch.TxHash.Hex())
						err := bc.Validator().ValidateLendingOrder(statedb, lendingState, tradingState, batch, author, block.Header())
						if err != nil {
							bc.reportBlock(block, nil, err)
							return i, events, coalescedLogs, err
						}
					}
					// liquidate / finalize open lendingTrades
					if block.Number().Uint64()%bc.chainConfig.XDPoS.Epoch == common.LiquidateLendingTradeBlock {
						finalizedTrades, _, _, _, _, err := lendingService.ProcessLiquidationData(block.Header(), bc, statedb, tradingState, lendingState)
						if err != nil {
							return i, events, coalescedLogs, fmt.Errorf("failed to ProcessLiquidationData. Err: %v", err)
						}
						if isSDKNode {
							finalizedTx := lendingstate.FinalizedResult{}
							if finalizedTx, err = ExtractLendingFinalizedTradeTransactions(block.Transactions()); err != nil {
								return i, events, coalescedLogs, err
							}
							bc.AddFinalizedTrades(finalizedTx.TxHash, finalizedTrades)
						}
					}
				}
				//check
				if tradingState != nil {
					gotRoot := tradingState.IntermediateRoot()
					expectRoot, _ := tradingService.GetTradingStateRoot(block, author)
					parentRoot, _ := tradingService.GetTradingStateRoot(parent, parentAuthor)
					if gotRoot != expectRoot {
						err = fmt.Errorf("invalid XDCx trading state merke trie got : %s , expect : %s ,parent : %s", gotRoot.Hex(), expectRoot.Hex(), parentRoot.Hex())
						bc.reportBlock(block, nil, err)
						return i, events, coalescedLogs, err
					}
					log.Debug("XDCX Trading State Root", "number", block.NumberU64(), "parent", parentRoot.Hex(), "nextRoot", expectRoot.Hex())
				}
				if lendingState != nil && tradingState != nil {
					gotRoot := lendingState.IntermediateRoot()
					expectRoot, _ := lendingService.GetLendingStateRoot(block, author)
					parentRoot, _ := lendingService.GetLendingStateRoot(parent, parentAuthor)
					if gotRoot != expectRoot {
						err = fmt.Errorf("invalid lending state merke trie got: %s, expect: %s, parent: %s", gotRoot.Hex(), expectRoot.Hex(), parentRoot.Hex())
						bc.reportBlock(block, nil, err)
						return i, events, coalescedLogs, err
					}
					log.Debug("XDCX Lending State Root", "number", block.NumberU64(), "parent", parentRoot.Hex(), "nextRoot", expectRoot.Hex())
				}
			}
		}
		feeCapacity := state.GetTRC21FeeCapacityFromStateWithCache(parent.Root(), statedb)
		// Process block using the parent state as reference point.
		t0 := time.Now()
		receipts, logs, usedGas, err := bc.processor.Process(block, statedb, tradingState, bc.vmConfig, feeCapacity)
		t1 := time.Now()
		if err != nil {
			bc.reportBlock(block, receipts, err)
			return i, events, coalescedLogs, err
		}
		// Validate the state using the default validator
		err = bc.Validator().ValidateState(block, parent, statedb, receipts, usedGas)
		if err != nil {
			bc.reportBlock(block, receipts, err)
			return i, events, coalescedLogs, err
		}
		t2 := time.Now()
		proctime := time.Since(bstart)

		// Write the block to the chain and get the status.
		status, err := bc.writeBlockWithState(block, receipts, statedb, tradingState, lendingState)
		t3 := time.Now()
		if err != nil {
			return i, events, coalescedLogs, err
		}

		// Update the metrics subsystem with all the measurements
		accountReadTimer.Update(statedb.AccountReads)
		accountHashTimer.Update(statedb.AccountHashes)
		accountUpdateTimer.Update(statedb.AccountUpdates)
		accountCommitTimer.Update(statedb.AccountCommits)

		storageReadTimer.Update(statedb.StorageReads)
		storageHashTimer.Update(statedb.StorageHashes)
		storageUpdateTimer.Update(statedb.StorageUpdates)
		storageCommitTimer.Update(statedb.StorageCommits)

		trieAccess := statedb.AccountReads + statedb.AccountHashes + statedb.AccountUpdates + statedb.AccountCommits
		trieAccess += statedb.StorageReads + statedb.StorageHashes + statedb.StorageUpdates + statedb.StorageCommits

		blockInsertTimer.UpdateSince(bstart)
		blockExecutionTimer.Update(t1.Sub(t0) - trieAccess)
		blockValidationTimer.Update(t2.Sub(t1))
		blockWriteTimer.Update(t3.Sub(t2))

		switch status {
		case CanonStatTy:
			log.Debug("Inserted new block from downloader", "number", block.Number(), "hash", block.Hash(), "uncles", len(block.Uncles()),
				"txs", len(block.Transactions()), "gas", block.GasUsed(), "elapsed", common.PrettyDuration(time.Since(bstart)))

			coalescedLogs = append(coalescedLogs, logs...)
			events = append(events, ChainEvent{block, block.Hash(), logs})
			lastCanon = block

			// Only count canonical blocks for GC processing time
			bc.gcproc += proctime
			bc.UpdateBlocksHashCache(block)
			if bc.chainConfig.IsTIPXDCX(block.Number()) && bc.chainConfig.XDPoS != nil && block.NumberU64() > bc.chainConfig.XDPoS.Epoch {
				bc.logExchangeData(block)
				bc.logLendingData(block)
			}
		case SideStatTy:
			log.Debug("Inserted forked block from downloader", "number", block.Number(), "hash", block.Hash(), "diff", block.Difficulty(), "elapsed",
				common.PrettyDuration(time.Since(bstart)), "txs", len(block.Transactions()), "gas", block.GasUsed(), "uncles", len(block.Uncles()))
			events = append(events, ChainSideEvent{block})
			bc.UpdateBlocksHashCache(block)
		}
		stats.processed++
		stats.usedGas += usedGas
		dirty, _ := bc.stateCache.TrieDB().Size()
		stats.report(chain, i, dirty)
		if bc.chainConfig.XDPoS != nil {
			// epoch block
			isEpochSwithBlock, _, err := engine.IsEpochSwitch(chain[i].Header())
			if err != nil {
				log.Error("[insertChain] Error while checking and notifying channel CheckpointCh if the incoming block is epoch switch block", "Hash", block.Hash(), "Number", block.Number())
				bc.reportBlock(block, nil, err)
			}
			if isEpochSwithBlock {
				CheckpointCh <- 1
			}
		}
	}
	// Append a single chain head event if we've progressed the chain
	if lastCanon != nil && bc.CurrentBlock().Hash() == lastCanon.Hash() {
		log.Debug("New ChainHeadEvent ", "number", lastCanon.NumberU64(), "hash", lastCanon.Hash())
		events = append(events, ChainHeadEvent{lastCanon})
	}
	return 0, events, coalescedLogs, nil
}

func (bc *BlockChain) InsertBlock(block *types.Block) error {
	events, logs, err := bc.insertBlock(block)
	bc.PostChainEvents(events, logs)
	return err
}

func (bc *BlockChain) PrepareBlock(block *types.Block) (err error) {
	defer log.Debug("Done prepare block ", "number", block.NumberU64(), "hash", block.Hash(), "validator", block.Header().Validator, "err", err)
	if _, check := bc.resultProcess.Get(block.Hash()); check {
		log.Debug("Stop prepare a block because the result cached", "number", block.NumberU64(), "hash", block.Hash(), "validator", block.Header().Validator)
		return nil
	}
	if _, check := bc.calculatingBlock.Get(block.Hash()); check {
		log.Debug("Stop prepare a block because inserting", "number", block.NumberU64(), "hash", block.Hash(), "validator", block.Header().Validator)
		return nil
	}
	err = bc.engine.VerifyHeader(bc, block.Header(), false)
	if err != nil {
		return err
	}
	result, err := bc.getResultBlock(block, false)
	if err == nil {
		bc.resultProcess.Add(block.Hash(), result)
		return nil
	} else if err == ErrKnownBlock {
		return nil
	} else if err == ErrStopPreparingBlock {
		log.Debug("Stop prepare a block because calculating", "number", block.NumberU64(), "hash", block.Hash(), "validator", block.Header().Validator)
		return nil
	}
	return err
}

func (bc *BlockChain) getResultBlock(block *types.Block, verifiedM2 bool) (*ResultProcessBlock, error) {
	var calculatedBlock *CalculatedBlock
	if verifiedM2 {
		if result, check := bc.resultProcess.Get(block.HashNoValidator()); check {
			log.Debug("Get result block from cache ", "number", block.NumberU64(), "hash", block.Hash(), "hash no validator", block.HashNoValidator())
			return result, nil
		}
		log.Debug("Not found cache prepare block ", "number", block.NumberU64(), "hash", block.Hash(), "validator", block.HashNoValidator())
		if calculatedBlock, _ := bc.calculatingBlock.Get(block.HashNoValidator()); calculatedBlock != nil {
			calculatedBlock.stop = true
		}
	}
	calculatedBlock = &CalculatedBlock{block, false}
	bc.calculatingBlock.Add(block.HashNoValidator(), calculatedBlock)
	// Start the parallel header verifier
	// If the chain is terminating, stop processing blocks
	if atomic.LoadInt32(&bc.procInterrupt) == 1 {
		log.Debug("Premature abort during blocks processing")
		return nil, ErrBlacklistedHash
	}
	// If the header is a banned one, straight out abort
	if BadHashes[block.Hash()] {
		bc.reportBlock(block, nil, ErrBlacklistedHash)
		return nil, ErrBlacklistedHash
	}
	// Wait for the block's verification to complete
	bstart := time.Now()
	err := bc.Validator().ValidateBody(block)
	switch {
	case err == ErrKnownBlock:
		// Block and state both already known. However if the current block is below
		// this number we did a rollback and we should reimport it nonetheless.
		if bc.CurrentBlock().NumberU64() >= block.NumberU64() {
			return nil, ErrKnownBlock
		}
	case err == consensus.ErrPrunedAncestor:
		// Block competing with the canonical chain, store in the db, but don't process
		// until the competitor TD goes above the canonical TD
		currentBlock := bc.CurrentBlock()
		localTd := bc.GetTd(currentBlock.Hash(), currentBlock.NumberU64())
		externTd := new(big.Int).Add(bc.GetTd(block.ParentHash(), block.NumberU64()-1), block.Difficulty())
		if localTd.Cmp(externTd) > 0 {
			return nil, err
		}
		// Competitor chain beat canonical, gather all blocks from the common ancestor
		var winner []*types.Block

		parent := bc.GetBlock(block.ParentHash(), block.NumberU64()-1)
		for !bc.HasFullState(parent) {
			winner = append(winner, parent)
			parent = bc.GetBlock(parent.ParentHash(), parent.NumberU64()-1)
		}
		for j := 0; j < len(winner)/2; j++ {
			winner[j], winner[len(winner)-1-j] = winner[len(winner)-1-j], winner[j]
		}
		log.Debug("Number block need calculated again", "number", block.NumberU64(), "hash", block.Hash().Hex(), "winners", len(winner))
		// Import all the pruned blocks to make the state available
		// During reorg, we use verifySeals=false
		_, _, _, err := bc.insertChain(winner, false)
		if err != nil {
			return nil, err
		}
	case err != nil:
		bc.reportBlock(block, nil, err)
		return nil, err
	}
	// Create a new statedb using the parent block and report an
	// error if it fails.
	var parent = bc.GetBlock(block.ParentHash(), block.NumberU64()-1)
	statedb, err := state.New(parent.Root(), bc.stateCache)
	if err != nil {
		return nil, err
	}
	engine, _ := bc.Engine().(*XDPoS.XDPoS)
	author, err := bc.Engine().Author(block.Header()) // Ignore error, we're past header validation
	if err != nil {
		bc.reportBlock(block, nil, err)
		return nil, err
	}
	parentAuthor, _ := bc.Engine().Author(parent.Header())

	var tradingState *tradingstate.TradingStateDB
	var lendingState *lendingstate.LendingStateDB
	var tradingService utils.TradingService
	var lendingService utils.LendingService
	isSDKNode := false
	if bc.Config().IsTIPXDCX(block.Number()) && bc.chainConfig.XDPoS != nil && engine != nil && block.NumberU64() > bc.chainConfig.XDPoS.Epoch {
		tradingService = engine.GetXDCXService()
		lendingService = engine.GetLendingService()
		if tradingService != nil && lendingService != nil {
			isSDKNode = tradingService.IsSDKNode()
			tradingState, err = tradingService.GetTradingState(parent, parentAuthor)
			if err != nil {
				bc.reportBlock(block, nil, err)
				return nil, err
			}
			lendingState, err = lendingService.GetLendingState(parent, parentAuthor)
			if err != nil {
				bc.reportBlock(block, nil, err)
				return nil, err
			}

			isEpochSwithBlock, epochNumber, err := engine.IsEpochSwitch(block.Header())
			if err != nil {
				log.Error("[getResultBlock] Error while checking block is epoch switch block", "Hash", block.Hash(), "Number", block.Number())
				bc.reportBlock(block, nil, err)
			}

			if isEpochSwithBlock {
				if err := tradingService.UpdateMediumPriceBeforeEpoch(epochNumber, tradingState, statedb); err != nil {
					return nil, err
				}
			} else {
				txMatchBatchData, err := ExtractTradingTransactions(block.Transactions())
				if err != nil {
					bc.reportBlock(block, nil, err)
					return nil, err
				}
				for _, txMatchBatch := range txMatchBatchData {
					log.Debug("Verify matching transaction", "txHash", txMatchBatch.TxHash.Hex())
					err := bc.Validator().ValidateTradingOrder(statedb, tradingState, txMatchBatch, author, block.Header())
					if err != nil {
						bc.reportBlock(block, nil, err)
						return nil, err
					}
				}
				batches, err := ExtractLendingTransactions(block.Transactions())
				if err != nil {
					bc.reportBlock(block, nil, err)
					return nil, err
				}
				for _, batch := range batches {
					log.Debug("Lending Verify matching transaction", "txHash", batch.TxHash.Hex())
					err := bc.Validator().ValidateLendingOrder(statedb, lendingState, tradingState, batch, author, block.Header())
					if err != nil {
						bc.reportBlock(block, nil, err)
						return nil, err
					}
				}
				// liquidate / finalize open lendingTrades
				if block.Number().Uint64()%bc.chainConfig.XDPoS.Epoch == common.LiquidateLendingTradeBlock {
					finalizedTrades, _, _, _, _, err := lendingService.ProcessLiquidationData(block.Header(), bc, statedb, tradingState, lendingState)
					if err != nil {
						return nil, fmt.Errorf("failed to ProcessLiquidationData. Err: %v", err)
					}
					if isSDKNode {
						finalizedTx := lendingstate.FinalizedResult{}
						if finalizedTx, err = ExtractLendingFinalizedTradeTransactions(block.Transactions()); err != nil {
							return nil, err
						}
						bc.AddFinalizedTrades(finalizedTx.TxHash, finalizedTrades)
					}
				}
			}
			if tradingState != nil {
				gotRoot := tradingState.IntermediateRoot()
				expectRoot, _ := tradingService.GetTradingStateRoot(block, author)
				parentRoot, _ := tradingService.GetTradingStateRoot(parent, parentAuthor)
				if gotRoot != expectRoot {
					err = fmt.Errorf("invalid XDCx trading state merke trie got : %s , expect : %s ,parent : %s", gotRoot.Hex(), expectRoot.Hex(), parentRoot.Hex())
					bc.reportBlock(block, nil, err)
					return nil, err
				}
				log.Debug("XDCX Trading State Root", "number", block.NumberU64(), "parent", parentRoot.Hex(), "nextRoot", expectRoot.Hex())
			}
			if lendingState != nil && tradingState != nil {
				gotRoot := lendingState.IntermediateRoot()
				expectRoot, _ := lendingService.GetLendingStateRoot(block, author)
				parentRoot, _ := lendingService.GetLendingStateRoot(parent, parentAuthor)
				if gotRoot != expectRoot {
					err = fmt.Errorf("invalid lending state merke trie got: %s , expect : %s , parent : %s", gotRoot.Hex(), expectRoot.Hex(), parentRoot.Hex())
					bc.reportBlock(block, nil, err)
					return nil, err
				}
				log.Debug("XDCX Lending State Root", "number", block.NumberU64(), "parent", parentRoot.Hex(), "nextRoot", expectRoot.Hex())
			}
		}
	}
	feeCapacity := state.GetTRC21FeeCapacityFromStateWithCache(parent.Root(), statedb)
	// Process block using the parent state as reference point.
	receipts, logs, usedGas, err := bc.processor.ProcessBlockNoValidator(calculatedBlock, statedb, tradingState, bc.vmConfig, feeCapacity)
	process := time.Since(bstart)
	if err != nil {
		if err != ErrStopPreparingBlock {
			bc.reportBlock(block, receipts, err)
		}
		return nil, err
	}
	// Validate the state using the default validator
	err = bc.Validator().ValidateState(block, parent, statedb, receipts, usedGas)
	if err != nil {
		bc.reportBlock(block, receipts, err)
		return nil, err
	}
	proctime := time.Since(bstart)
	log.Debug("Calculate new block", "number", block.Number(), "hash", block.Hash(), "uncles", len(block.Uncles()),
		"txs", len(block.Transactions()), "gas", block.GasUsed(), "elapsed", common.PrettyDuration(time.Since(bstart)), "process", process)
	return &ResultProcessBlock{receipts: receipts, logs: logs, state: statedb, tradingState: tradingState, lendingState: lendingState, proctime: proctime, usedGas: usedGas}, nil
}

// UpdateBlocksHashCache update BlocksHashCache by block number
func (bc *BlockChain) UpdateBlocksHashCache(block *types.Block) []common.Hash {
	var hashArr []common.Hash
	blockNumber := block.Number().Uint64()
	cached, ok := bc.blocksHashCache.Get(blockNumber)

	if ok {
		hashArr := cached
		hashArr = append(hashArr, block.Hash())
		bc.blocksHashCache.Remove(blockNumber)
		bc.blocksHashCache.Add(blockNumber, hashArr)
		return hashArr
	}

	hashArr = []common.Hash{
		block.Hash(),
	}
	bc.blocksHashCache.Add(blockNumber, hashArr)
	return hashArr
}

// insertChain will execute the actual chain insertion and event aggregation. The
// only reason this method exists as a separate one is to make locking cleaner
// with deferred statements.
func (bc *BlockChain) insertBlock(block *types.Block) ([]interface{}, []*types.Log, error) {
	var (
		stats         = insertStats{startTime: mclock.Now()}
		events        = make([]interface{}, 0, 1)
		coalescedLogs []*types.Log
	)
	if _, check := bc.downloadingBlock.Get(block.Hash()); check {
		log.Debug("Stop fetcher a block because downloading", "number", block.NumberU64(), "hash", block.Hash())
		return events, coalescedLogs, nil
	}
	result, err := bc.getResultBlock(block, true)
	if err != nil {
		return events, coalescedLogs, err
	}
	defer bc.resultProcess.Remove(block.HashNoValidator())
	bc.wg.Add(1)
	defer bc.wg.Done()
	// Write the block to the chain and get the status.
	if !bc.chainmu.TryLock() {
		return nil, nil, errChainStopped
	}
	defer bc.chainmu.Unlock()
	if bc.HasBlockAndFullState(block.Hash(), block.NumberU64()) {
		return events, coalescedLogs, nil
	}
	status, err := bc.writeBlockWithState(block, result.receipts, result.state, result.tradingState, result.lendingState)

	if err != nil {
		return events, coalescedLogs, err
	}
	switch status {
	case CanonStatTy:
		log.Debug("Inserted new block from fetcher", "number", block.Number(), "hash", block.Hash(), "uncles", len(block.Uncles()),
			"txs", len(block.Transactions()), "gas", block.GasUsed(), "elapsed", common.PrettyDuration(time.Since(block.ReceivedAt)))
		coalescedLogs = append(coalescedLogs, result.logs...)
		events = append(events, ChainEvent{block, block.Hash(), result.logs})
		// Only count canonical blocks for GC processing time
		bc.gcproc += result.proctime
		bc.UpdateBlocksHashCache(block)
		if bc.chainConfig.IsTIPXDCXReceiver(block.Number()) && bc.chainConfig.XDPoS != nil && block.NumberU64() > bc.chainConfig.XDPoS.Epoch {
			bc.logExchangeData(block)
			bc.logLendingData(block)
		}
	case SideStatTy:
		log.Debug("Inserted forked block from fetcher", "number", block.Number(), "hash", block.Hash(), "diff", block.Difficulty(), "elapsed",
			common.PrettyDuration(time.Since(block.ReceivedAt)), "txs", len(block.Transactions()), "gas", block.GasUsed(), "uncles", len(block.Uncles()))
		blockInsertTimer.Update(result.proctime)
		events = append(events, ChainSideEvent{block})
		bc.UpdateBlocksHashCache(block)
	}
	stats.processed++
	stats.usedGas += result.usedGas
	dirty, _ := bc.stateCache.TrieDB().Size()
	stats.report(types.Blocks{block}, 0, dirty)
	if bc.chainConfig.XDPoS != nil {
		// epoch block
		isEpochSwithBlock, _, err := bc.Engine().(*XDPoS.XDPoS).IsEpochSwitch(block.Header())
		if err != nil {
			log.Error("[insertBlock] Error while checking if the incoming block is epoch switch block", "Hash", block.Hash(), "Number", block.Number())
			bc.reportBlock(block, nil, err)
		}
		if isEpochSwithBlock {
			CheckpointCh <- 1
		}
	}
	// Append a single chain head event if we've progressed the chain
	if status == CanonStatTy && bc.CurrentBlock().Hash() == block.Hash() {
		events = append(events, ChainHeadEvent{block})
		log.Debug("New ChainHeadEvent from fetcher ", "number", block.NumberU64(), "hash", block.Hash())
	}
	return events, coalescedLogs, nil
}

// insertStats tracks and reports on block insertion.
type insertStats struct {
	queued, processed, ignored int
	usedGas                    uint64
	lastIndex                  int
	startTime                  mclock.AbsTime
}

// statsReportLimit is the time limit during import after which we always print
// out progress. This avoids the user wondering what's going on.
const statsReportLimit = 8 * time.Second

// report prints statistics if some number of blocks have been processed
// or more than a few seconds have passed since the last message.
func (st *insertStats) report(chain []*types.Block, index int, dirty common.StorageSize) {
	// Fetch the timings for the batch
	var (
		now     = mclock.Now()
		elapsed = time.Duration(now) - time.Duration(st.startTime)
	)
	// If we're at the last block of the batch or report period reached, log
	if index == len(chain)-1 || elapsed >= statsReportLimit {
		var (
			end = chain[index]
			txs = countTransactions(chain[st.lastIndex : index+1])
		)
		context := []interface{}{
			"blocks", st.processed, "txs", txs, "mgas", float64(st.usedGas) / 1000000,
			"elapsed", common.PrettyDuration(elapsed), "mgasps", float64(st.usedGas) * 1000 / float64(elapsed),
			"number", end.Number(), "hash", end.Hash(), "dirty", dirty,
		}
		if st.queued > 0 {
			context = append(context, []interface{}{"queued", st.queued}...)
		}
		if st.ignored > 0 {
			context = append(context, []interface{}{"ignored", st.ignored}...)
		}
		log.Info("Imported new chain segment", context...)
		*st = insertStats{startTime: now, lastIndex: index + 1}
	}
}

func countTransactions(chain []*types.Block) (c int) {
	for _, b := range chain {
		c += len(b.Transactions())
	}
	return c
}

// collectLogs collects the logs that were generated or removed during
// the processing of a block. These logs are later announced as deleted or reborn.
func (bc *BlockChain) collectLogs(b *types.Block, removed bool) []*types.Log {
	receipts := rawdb.ReadRawReceipts(bc.db, b.Hash(), b.NumberU64())
	if err := receipts.DeriveFields(bc.chainConfig, b.Hash(), b.NumberU64(), b.BaseFee(), b.Transactions()); err != nil {
		log.Error("Failed to derive block receipts fields", "hash", b.Hash(), "number", b.NumberU64(), "err", err)
	}

	var logs []*types.Log
	for _, receipt := range receipts {
		for _, log := range receipt.Logs {
			if removed {
				log.Removed = true
			}
			logs = append(logs, log)
		}
	}
	return logs
}

// reorg takes two blocks, an old chain and a new chain and will reconstruct the
// blocks and inserts them to be part of the new canonical chain and accumulates
// potential missing transactions and post an event about them.
func (bc *BlockChain) reorg(oldHead, newHead *types.Header) error {
	log.Warn("Reorg", "OldHash", oldHead.Hash().Hex(), "OldNum", oldHead.Number, "NewHash", newHead.Hash().Hex(), "NewNum", newHead.Number)

	var (
		newChain    []*types.Header
		oldChain    []*types.Header
		commonBlock *types.Header
	)

	// Reduce the longer chain to the same number as the shorter one
	if oldHead.Number.Uint64() > newHead.Number.Uint64() {
		// Old chain is longer, gather all transactions and logs as deleted ones
		for ; oldHead != nil && oldHead.Number.Uint64() != newHead.Number.Uint64(); oldHead = bc.GetHeader(oldHead.ParentHash, oldHead.Number.Uint64()-1) {
			oldChain = append(oldChain, oldHead)
		}
	} else {
		// New chain is longer, stash all blocks away for subsequent insertion
		for ; newHead != nil && newHead.Number.Uint64() != oldHead.Number.Uint64(); newHead = bc.GetHeader(newHead.ParentHash, newHead.Number.Uint64()-1) {
			newChain = append(newChain, newHead)
		}
	}
	if oldHead == nil {
		return errInvalidOldChain
	}
	if newHead == nil {
		return errInvalidNewChain
	}

	// Both sides of the reorg are at the same number, reduce both until the common
	// ancestor is found
	for {
		// If the common ancestor was found, bail out
		if oldHead.Hash() == newHead.Hash() {
			commonBlock = oldHead
			break
		}
		// Remove an old block as well as stash away a new block
		oldChain = append(oldChain, oldHead)
		newChain = append(newChain, newHead)

		// Step back with both chains
		oldHead = bc.GetHeader(oldHead.ParentHash, oldHead.Number.Uint64()-1)
		if oldHead == nil {
			return errInvalidOldChain
		}
		newHead = bc.GetHeader(newHead.ParentHash, newHead.Number.Uint64()-1)
		if newHead == nil {
			return errInvalidNewChain
		}
	}

	// Ensure XDPoS engine committed block will be not reverted
	if xdpos, ok := bc.Engine().(*XDPoS.XDPoS); ok {
		latestCommittedBlock := xdpos.EngineV2.GetLatestCommittedBlockInfo()
		if latestCommittedBlock != nil {
			cmp := commonBlock.Number.Cmp(latestCommittedBlock.Number)
			if cmp < 0 {
				for _, oldBlock := range oldChain {
					if oldBlock.Number.Cmp(latestCommittedBlock.Number) == 0 {
						if oldBlock.Hash() != latestCommittedBlock.Hash {
							log.Error("Impossible reorg, please file an issue", "OldNum", oldBlock.Number, "OldHash", oldBlock.Hash().Hex(), "LatestCommittedHash", latestCommittedBlock.Hash.Hex())
						} else {
							log.Warn("Stop reorg, blockchain is under forking attack", "OldCommittedNum", oldBlock.Number, "OldCommittedHash", oldBlock.Hash().Hex())
							return fmt.Errorf("stop reorg, blockchain is under forking attack. OldCommitted num %d, hash %s", oldBlock.Number, oldBlock.Hash().Hex())
						}
					}
				}
			} else if cmp == 0 {
				if commonBlock.Hash() != latestCommittedBlock.Hash {
					log.Error("Impossible reorg, please file an issue", "OldNum", commonBlock.Number.Uint64(), "OldHash", commonBlock.Hash().Hex(), "LatestCommittedHash", latestCommittedBlock.Hash.Hex())
				}
			}
		}
	}

	// Ensure the user sees large reorgs
	if len(oldChain) > 0 && len(newChain) > 0 {
		logFn := log.Info
		msg := "Chain reorg detected"
		if len(oldChain) > 63 {
			msg = "Large chain reorg detected"
			logFn = log.Warn
		}
		logFn(msg, "number", commonBlock.Number, "hash", commonBlock.Hash().Hex(),
			"drop", len(oldChain), "dropfrom", oldChain[0].Hash().Hex(), "add", len(newChain), "addfrom", newChain[0].Hash().Hex())
		blockReorgAddMeter.Mark(int64(len(newChain)))
		blockReorgDropMeter.Mark(int64(len(oldChain)))
		blockReorgMeter.Mark(1)
	} else if len(newChain) > 0 {
		// Special case happens in the post merge stage that current head is
		// the ancestor of new head while these two blocks are not consecutive
		log.Info("Extend chain", "add", len(newChain), "number", newChain[0].Number, "hash", newChain[0].Hash())
		blockReorgAddMeter.Mark(int64(len(newChain)))
	} else {
		// len(newChain) == 0 && len(oldChain) > 0
		// rewind the canonical chain to a lower point.
		log.Error("Impossible reorg, please file an issue", "oldnum", oldHead.Number, "oldhash", oldHead.Hash(), "oldblocks", len(oldChain), "newnum", newHead.Number, "newhash", newHead.Hash(), "newblocks", len(newChain))
	}

	// Acquire the tx-lookup lock before mutation. This step is essential
	// as the txlookups should be changed atomically, and all subsequent
	// reads should be blocked until the mutation is complete.
	// bc.txLookupLock.Lock()

	// Reorg can be executed, start reducing the chain's old blocks and appending
	// the new blocks
	var (
		deletedTxs []common.Hash
		rebirthTxs []common.Hash

		deletedLogs []*types.Log
		rebirthLogs []*types.Log
	)

	// Deleted log emission on the API uses forward order, which is borked, but
	// we'll leave it in for legacy reasons.
	//
	// TODO(karalabe): This should be nuked out, no idea how, deprecate some APIs?
	{
		for i := len(oldChain) - 1; i >= 0; i-- {
			block := bc.GetBlock(oldChain[i].Hash(), oldChain[i].Number.Uint64())
			if block == nil {
				return errInvalidOldChain // Corrupt database, mostly here to avoid weird panics
			}
			if logs := bc.collectLogs(block, true); len(logs) > 0 {
				deletedLogs = append(deletedLogs, logs...)
			}
			if len(deletedLogs) > 512 {
				go bc.rmLogsFeed.Send(RemovedLogsEvent{deletedLogs})
				deletedLogs = nil
			}
			// TODO(daniel): remove chainSideFeed, reference PR #30601
			// Also send event for blocks removed from the canon chain.
			// bc.chainSideFeed.Send(ChainSideEvent{Block: block})
		}
		if len(deletedLogs) > 0 {
			go bc.rmLogsFeed.Send(RemovedLogsEvent{deletedLogs})
		}
	}

	// Undo old blocks in reverse order
	for i := 0; i < len(oldChain); i++ {
		// Collect all the deleted transactions
		block := bc.GetBlock(oldChain[i].Hash(), oldChain[i].Number.Uint64())
		if block == nil {
			return errInvalidOldChain // Corrupt database, mostly here to avoid weird panics
		}
		for _, tx := range block.Transactions() {
			deletedTxs = append(deletedTxs, tx.Hash())
		}
		// Collect deleted logs and emit them for new integrations
		// if logs := bc.collectLogs(block, true); len(logs) > 0 {
		// 	slices.Reverse(logs) // Emit revertals latest first, older then
		// }
	}

	// Apply new blocks in forward order
	for i := len(newChain) - 1; i >= 0; i-- {
		// Collect all the included transactions
		block := bc.GetBlock(newChain[i].Hash(), newChain[i].Number.Uint64())
		if block == nil {
			return errInvalidNewChain // Corrupt database, mostly here to avoid weird panics
		}
		for _, tx := range block.Transactions() {
			rebirthTxs = append(rebirthTxs, tx.Hash())
		}
		// Collect inserted logs and emit them
		if logs := bc.collectLogs(block, false); len(logs) > 0 {
			rebirthLogs = append(rebirthLogs, logs...)
		}
		if len(rebirthLogs) > 512 {
			bc.logsFeed.Send(rebirthLogs)
			rebirthLogs = nil
		}
		// Update the head block
		bc.writeHeadBlock(block, true)
		// prepare set of masternodes for the next epoch
		if bc.chainConfig.XDPoS != nil && ((block.NumberU64() % bc.chainConfig.XDPoS.Epoch) == (bc.chainConfig.XDPoS.Epoch - bc.chainConfig.XDPoS.Gap)) {
			if err := bc.UpdateM1(); err != nil {
				log.Crit("Fail to update masternodes during reorg", "number", block.Number, "hash", block.Hash().Hex(), "err", err)
			}
		}
	}
	if len(rebirthLogs) > 0 {
		bc.logsFeed.Send(rebirthLogs)
	}

	// Delete useless indexes right now which includes the non-canonical
	// transaction indexes, canonical chain indexes which above the head.
	batch := bc.db.NewBatch()
	for _, tx := range types.HashDifference(deletedTxs, rebirthTxs) {
		rawdb.DeleteTxLookupEntry(batch, tx)
	}
	// Delete all hash markers that are not part of the new canonical chain.
	// Because the reorg function handles new chain head, all hash
	// markers greater than new chain head should be deleted.
	number := commonBlock.Number
	if len(newChain) > 0 {
		number = newChain[0].Number
	}
	for i := number.Uint64() + 1; ; i++ {
		hash := rawdb.ReadCanonicalHash(bc.db, i)
		if hash == (common.Hash{}) {
			break
		}
		rawdb.DeleteCanonicalHash(batch, i)
	}
	if err := batch.Write(); err != nil {
		log.Crit("Failed to delete useless indexes", "err", err)
	}

	// Reset the tx lookup cache to clear stale txlookup cache.
	// bc.txLookupCache.Purge()

	// Release the tx-lookup lock after mutation.
	// bc.txLookupLock.Unlock()

	return nil
}

// PostChainEvents iterates over the events generated by a chain insertion and
// posts them into the event feed.
// TODO: Should not expose PostChainEvents. The chain events should be posted in WriteBlock.
func (bc *BlockChain) PostChainEvents(events []interface{}, logs []*types.Log) {
	// post event logs for further processing
	if logs != nil {
		bc.logsFeed.Send(logs)
	}
	for _, event := range events {
		switch ev := event.(type) {
		case ChainEvent:
			bc.chainFeed.Send(ev)

		case ChainHeadEvent:
			bc.chainHeadFeed.Send(ev)

		case ChainSideEvent:
			bc.chainSideFeed.Send(ev)
		}
	}
}

// futureBlocksLoop processes the 'future block' queue.
func (bc *BlockChain) futureBlocksLoop() {
	defer bc.wg.Done()

	futureTimer := time.NewTicker(10 * time.Millisecond)
	defer futureTimer.Stop()
	for {
		select {
		case <-futureTimer.C:
			bc.procFutureBlocks()
		case <-bc.quit:
			return
		}
	}
}

// BadBlockArgs represents the entries in the list returned when bad blocks are queried.
type BadBlockArgs struct {
	Hash   common.Hash   `json:"hash"`
	Header *types.Header `json:"header"`
}

// BadBlocks returns a list of the last 'bad blocks' that the client has seen on the network
func (bc *BlockChain) BadBlocks() ([]BadBlockArgs, error) {
	headers := make([]BadBlockArgs, 0, bc.badBlocks.Len())
	for _, hash := range bc.badBlocks.Keys() {
		if header, exist := bc.badBlocks.Peek(hash); exist {
			headers = append(headers, BadBlockArgs{header.Hash(), header})
		}
	}
	return headers, nil
}

// addBadBlock adds a bad block to the bad-block LRU cache
func (bc *BlockChain) addBadBlock(block *types.Block) {
	bc.badBlocks.Add(block.Header().Hash(), block.Header())
}

// reportBlock logs a bad block error.
func (bc *BlockChain) reportBlock(block *types.Block, receipts types.Receipts, err error) {
	bc.addBadBlock(block)

	var roundNumber = types.Round(0)
	engine, ok := bc.Engine().(*XDPoS.XDPoS)
	if ok {
		var err error
		roundNumber, err = engine.EngineV2.GetRoundNumber(block.Header())
		if err != nil {
			log.Error("reportBlock", "GetRoundNumber", err)
		}
	}

	var receiptString string
	for i, receipt := range receipts {
		receiptString += fmt.Sprintf("\n  %d: cumulative: %v gas: %v contract: %v status: %v tx: %v logs: %v bloom: %x state: %x",
			i, receipt.CumulativeGasUsed, receipt.GasUsed, receipt.ContractAddress.Hex(),
			receipt.Status, receipt.TxHash.Hex(), receipt.Logs, receipt.Bloom, receipt.PostState)
	}
	log.Error(fmt.Sprintf(`
########## BAD BLOCK #########
Number: %v
Hash: %#x
Round: %v
Error: %v
Chain config: %v
Receipts: %v
##############################
`, block.Number(), block.Hash(), roundNumber, err, bc.chainConfig, receiptString))
}

// InsertHeaderChain attempts to insert the given header chain in to the local
// chain, possibly creating a reorg. If an error is returned, it will return the
// index number of the failing header as well an error describing what went wrong.
//
// The verify parameter can be used to fine tune whether nonce verification
// should be done or not. The reason behind the optional check is because some
// of the header retrieval mechanisms already need to verify nonces, as well as
// because nonces can be verified sparsely, not needing to check each.
func (bc *BlockChain) InsertHeaderChain(chain []*types.Header, checkFreq int) (int, error) {
	start := time.Now()
	if i, err := bc.hc.ValidateHeaderChain(chain, checkFreq); err != nil {
		return i, err
	}

	if !bc.chainmu.TryLock() {
		return 0, errChainStopped
	}
	defer bc.chainmu.Unlock()

	whFunc := func(header *types.Header) error {
		_, err := bc.hc.WriteHeader(header)
		return err
	}

	return bc.hc.InsertHeaderChain(chain, whFunc, start)
}

// CurrentHeader retrieves the current head header of the canonical chain. The
// header is retrieved from the HeaderChain's internal cache.
func (bc *BlockChain) CurrentHeader() *types.Header {
	return bc.hc.CurrentHeader()
}

// GetTd retrieves a block's total difficulty in the canonical chain from the
// database by hash and number, caching it if found.
func (bc *BlockChain) GetTd(hash common.Hash, number uint64) *big.Int {
	return bc.hc.GetTd(hash, number)
}

// GetTdByHash retrieves a block's total difficulty in the canonical chain from the
// database by hash, caching it if found.
func (bc *BlockChain) GetTdByHash(hash common.Hash) *big.Int {
	return bc.hc.GetTdByHash(hash)
}

// GetHeader retrieves a block header from the database by hash and number,
// caching it if found.
func (bc *BlockChain) GetHeader(hash common.Hash, number uint64) *types.Header {
	return bc.hc.GetHeader(hash, number)
}

// GetHeaderByHash retrieves a block header from the database by hash, caching it if
// found.
func (bc *BlockChain) GetHeaderByHash(hash common.Hash) *types.Header {
	return bc.hc.GetHeaderByHash(hash)
}

// HasHeader checks if a block header is present in the database or not, caching
// it if present.
func (bc *BlockChain) HasHeader(hash common.Hash, number uint64) bool {
	return bc.hc.HasHeader(hash, number)
}

// GetCanonicalHash returns the canonical hash for a given block number
func (bc *BlockChain) GetCanonicalHash(number uint64) common.Hash {
	return bc.hc.GetCanonicalHash(number)
}

// GetBlockHashesFromHash retrieves a number of block hashes starting at a given
// hash, fetching towards the genesis block.
func (bc *BlockChain) GetBlockHashesFromHash(hash common.Hash, max uint64) []common.Hash {
	return bc.hc.GetBlockHashesFromHash(hash, max)
}

// GetHeaderByNumber retrieves a block header from the database by number,
// caching it (associated with its hash) if found.
func (bc *BlockChain) GetHeaderByNumber(number uint64) *types.Header {
	return bc.hc.GetHeaderByNumber(number)
}

// Set config for testing purpose function
func (bc *BlockChain) SetConfig(config *params.ChainConfig) {
	bc.chainConfig = config
}

// Config retrieves the blockchain's chain configuration.
func (bc *BlockChain) Config() *params.ChainConfig { return bc.chainConfig }

// Engine retrieves the blockchain's consensus engine.
func (bc *BlockChain) Engine() consensus.Engine { return bc.engine }

// SubscribeRemovedLogsEvent registers a subscription of RemovedLogsEvent.
func (bc *BlockChain) SubscribeRemovedLogsEvent(ch chan<- RemovedLogsEvent) event.Subscription {
	return bc.scope.Track(bc.rmLogsFeed.Subscribe(ch))
}

// SubscribeChainEvent registers a subscription of ChainEvent.
func (bc *BlockChain) SubscribeChainEvent(ch chan<- ChainEvent) event.Subscription {
	return bc.scope.Track(bc.chainFeed.Subscribe(ch))
}

// SubscribeChainHeadEvent registers a subscription of ChainHeadEvent.
func (bc *BlockChain) SubscribeChainHeadEvent(ch chan<- ChainHeadEvent) event.Subscription {
	return bc.scope.Track(bc.chainHeadFeed.Subscribe(ch))
}

// SubscribeChainSideEvent registers a subscription of ChainSideEvent.
func (bc *BlockChain) SubscribeChainSideEvent(ch chan<- ChainSideEvent) event.Subscription {
	return bc.scope.Track(bc.chainSideFeed.Subscribe(ch))
}

// SubscribeLogsEvent registers a subscription of []*types.Log.
func (bc *BlockChain) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return bc.scope.Track(bc.logsFeed.Subscribe(ch))
}

// Get current IPC Client.
func (bc *BlockChain) GetClient() (bind.ContractBackend, error) {
	if bc.Client == nil {
		// Inject ipc client global instance.
		client, err := ethclient.Dial(bc.IPCEndpoint)
		if err != nil {
			log.Error("Fail to connect IPC", "error", err)
			return nil, err
		}
		bc.Client = client
	}

	return bc.Client, nil
}

func (bc *BlockChain) UpdateM1() error {
	engine, ok := bc.Engine().(*XDPoS.XDPoS)
	if bc.Config().XDPoS == nil || !ok {
		return ErrNotXDPoS
	}
	log.Info("It's time to update new set of masternodes for the next epoch...")
	// get masternodes information from smart contract
	client, err := bc.GetClient()
	if err != nil {
		return err
	}
	addr := common.MasternodeVotingSMCBinary
	validator, err := contractValidator.NewXDCValidator(addr, client)
	if err != nil {
		return err
	}
	opts := new(bind.CallOpts)

	var candidates []common.Address
	// get candidates from slot of stateDB
	// if can't get anything, request from contracts
	stateDB, err := bc.State()
	if err != nil {
		candidates, err = validator.GetCandidates(opts)
		if err != nil {
			return err
		}
	} else if stateDB == nil {
		return errors.New("nil stateDB in UpdateM1")
	} else {
		candidates = state.GetCandidates(stateDB)
	}

	var ms []utils.Masternode
	for _, candidate := range candidates {
		v, err := validator.GetCandidateCap(opts, candidate)
		if err != nil {
			return err
		}
		// TODO: smart contract shouldn't return "0x0000000000000000000000000000000000000000"
		if !candidate.IsZero() {
			ms = append(ms, utils.Masternode{Address: candidate, Stake: v})
		}
	}
	if len(ms) == 0 {
		log.Error("No masternode found. Stopping node")
		os.Exit(1)
	} else {
		sort.Slice(ms, func(i, j int) bool {
			return ms[i].Stake.Cmp(ms[j].Stake) >= 0
		})
		log.Info("Ordered list of masternode candidates")
		for _, m := range ms {
			log.Info("", "address", m.Address.String(), "stake", m.Stake)
		}
		// update masternodes

		log.Info("Updating new set of masternodes")
		// get block header
		header := bc.CurrentHeader()
		err = engine.UpdateMasternodes(bc, header, ms)
		if err != nil {
			return err
		}
		log.Info("Masternodes are ready for the next epoch")
	}
	return nil
}

// postChainEvents iterates over the events generated by a chain insertion and
// posts them into the event mux.
func (self *BlockChain) postChainEvents(events []interface{}) {
	for _, event := range events {
		if event, ok := event.(ChainEvent); ok {
			// We need some control over the mining operation. Acquiring locks and waiting for the miner to create new block takes too long
			// and in most cases isn't even necessary.
			if self.LastBlockHash() == event.Hash {
				self.eventMux.Post(ChainHeadEvent{event.Block})
			}
		}
		// Fire the insertion events individually too
		self.eventMux.Post(event)
	}
}

func (self *BlockChain) update() {
	futureTimer := time.Tick(5 * time.Second)
	for {
		select {
		case <-futureTimer:
			self.procFutureBlocks()
		case <-self.quit:
			return
		}
	}
}

func blockErr(block *types.Block, err error) {
	if glog.V(logger.Error) {
		glog.Errorf("Bad block #%v (%s)\n", block.Number(), block.Hash().Hex())
		glog.Errorf("    %v", err)
	}
}

func (bc *BlockChain) AddFinalizedTrades(txHash common.Hash, trades map[common.Hash]*lendingstate.LendingTrade) {
	bc.finalizedTrade.Add(txHash, trades)
}
