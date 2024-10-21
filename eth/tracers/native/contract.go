// Copyright 2021 The go-ethereum Authors
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

package native

import (
	"encoding/json"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/core/vm"
	"github.com/XinFinOrg/XDPoSChain/eth/tracers"
)

func init() {
	tracers.RegisterNativeTracer("contractTracer", NewContractTracer)
}

type contractTracer struct {
	Addrs     map[string]string
	config    contractTracerConfig
	interrupt uint32 // Atomic flag to signal execution interruption
	reason    error  // Textual reason for the interruption
}

type contractTracerConfig struct {
	OpCode       string `json:"opCode"`       // Target opcode to trace
	WithByteCode bool   `json:"withByteCode"` // If true, bytecode will be collected
}

// NewContractTracer returns a native go tracer which tracks the contracr was created
func NewContractTracer(cfg json.RawMessage) (tracers.Tracer, error) {
	var config contractTracerConfig
	if cfg != nil {
		if err := json.Unmarshal(cfg, &config); err != nil {
			return nil, err
		}
	}
	t := &contractTracer{
		Addrs:  make(map[string]string, 1),
		config: config,
	}
	// handle invalid opcode case
	op := vm.StringToOp(t.config.OpCode)
	if op == 0 && t.config.OpCode != "STOP" && t.config.OpCode != "" {
		t.config.OpCode = "inv"
	}
	return t, nil
}

func (t *contractTracer) CaptureStart(env *vm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) {
	if create {
		validateAndStoreOpCode(t, input, to)
	}
}

func (t *contractTracer) CaptureEnd(output []byte, gasUsed uint64, _ time.Duration, err error) {
}

func (t *contractTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {
}

func (t *contractTracer) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, _ *vm.ScopeContext, depth int, err error) {
}

func (t *contractTracer) CaptureEnter(typ vm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
	// Skip if tracing was interrupted
	if atomic.LoadUint32(&t.interrupt) > 0 {
		// TODO: env.Cancel()
		return
	}
	if typ == vm.CREATE || typ == vm.CREATE2 {
		validateAndStoreOpCode(t, input, to)
	}
}

func (t *contractTracer) CaptureExit(output []byte, gasUsed uint64, err error) {
}

func (t *contractTracer) GetResult() (json.RawMessage, error) {
	res, err := json.Marshal(t.Addrs)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(res), t.reason
}

func (t *contractTracer) Stop(err error) {
	t.reason = err
	atomic.StoreUint32(&t.interrupt, 1)
}

func validateAndStoreOpCode(t *contractTracer, input []byte, to common.Address) {
	// If the OpCode is "inv" or if the OpCode is not empty and doesn't match the input, exit early.
	if t.config.OpCode == "inv" || (t.config.OpCode != "" && !findOpcodes(input, vm.StringToOp(t.config.OpCode))) {
		return
	}
	// If WithByteCode is true, store the input in the address mapping as hex.
	if t.config.WithByteCode {
		t.Addrs[addrToHex(to)] = bytesToHex(input)
	} else {
		t.Addrs[addrToHex(to)] = ""
	}
}

// Compare bytecode with the given opcode, skipping PUSH instructions.
func findOpcodes(bytecode []byte, opcode vm.OpCode) bool {
	for i := 0; i < len(bytecode); {
		op := vm.OpCode(bytecode[i])
		// Skip PUSH opcodes and their arguments
		if op.IsPush() {
			i += int(op - 95) // Directly calculate the number of bytes to skip
		} else if op == opcode {
			return true
		}
		i++
	}
	return false
}
