// Copyright 2019 The go-ethereum Authors
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

package backends

import (
	"context"
	"errors"
	"strings"

	"math/big"

	"testing"

	"github.com/XinFinOrg/XDPoSChain"
	"github.com/XinFinOrg/XDPoSChain/accounts/abi"
	"github.com/XinFinOrg/XDPoSChain/accounts/abi/bind"
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/core"
	"github.com/XinFinOrg/XDPoSChain/crypto"
	"github.com/XinFinOrg/XDPoSChain/params"
)

func TestSimulatedBackend_EstimateGas(t *testing.T) {
	/*
		pragma solidity ^0.6.4;
		contract GasEstimation {
		    function PureRevert() public { revert(); }
		    function Revert() public { revert("revert reason");}
		    function OOG() public { for (uint i = 0; ; i++) {}}
		    function Assert() public { assert(false);}
		    function Valid() public {}
		}*/
	const contractAbi = "[{\"inputs\":[],\"name\":\"Assert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OOG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PureRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Valid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	const contractBin = "0x608060405234801561001057600080fd5b50610156806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806350f6fe341461005c578063aa8b1d3014610066578063b9b046f914610070578063d8b983911461007a578063e09fface14610084575b600080fd5b61006461008e565b005b61006e6100a1565b005b6100786100a6565b005b6100826100b0565b005b61008c61011e565b005b60008090505b8080600101915050610094565b600080fd5b60006100ae57fe5b565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f72657665727420726561736f6e0000000000000000000000000000000000000081525060200191505060405180910390fd5b56fea26469706673582212206f8c043de30823c47c0df44a4404868a45bd4fbb4ff8846a6e1f476d79d3297764736f6c63430006040033"

	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	opts := bind.NewKeyedTransactor(key)

	sim := NewXDCSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(params.Ether)}}, 10000000, params.TestXDPoSMockChainConfig)
	defer sim.Close()

	parsed, _ := abi.JSON(strings.NewReader(contractAbi))
	contractAddr, _, _, _ := bind.DeployContract(opts, parsed, common.FromHex(contractBin), sim)
	sim.Commit()

	var cases = []struct {
		name        string
		message     XDPoSChain.CallMsg
		expect      uint64
		expectError error
	}{
		{"plain transfer(valid)", XDPoSChain.CallMsg{
			From:     addr,
			To:       &addr,
			Gas:      0,
			GasPrice: big.NewInt(0),
			Value:    big.NewInt(1),
			Data:     nil,
		}, params.TxGas, nil},

		{"plain transfer(invalid)", XDPoSChain.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      0,
			GasPrice: big.NewInt(0),
			Value:    big.NewInt(1),
			Data:     nil,
		}, 0, errors.New("always failing transaction (execution reverted)")},

		{"Revert", XDPoSChain.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      0,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("d8b98391"), //("d8b98391"),
		}, 0, errors.New("always failing transaction (execution reverted) (revert reason)")},

		{"PureRevert", XDPoSChain.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      0,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("aa8b1d30"), //aa8b1d30
		}, 0, errors.New("always failing transaction (execution reverted)")},

		{"OOG", XDPoSChain.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      100000,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("50f6fe34"),
		}, 0, errors.New("gas required exceeds allowance (100000)")},

		{"Assert", XDPoSChain.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      100000,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("b9b046f9"),
		}, 0, errors.New("always failing transaction (invalid opcode: opcode 0xfe not defined)")}, //opcode 0xfe not defined have not change

		{"Valid", XDPoSChain.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      100000,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("e09fface"),
		}, 21275, nil},
	}
	for _, c := range cases {
		got, err := sim.EstimateGas(context.Background(), c.message)
		if c.expectError != nil {
			if err == nil {
				t.Fatalf("Expect error, got nil")
			}
			if c.expectError.Error() != err.Error() {
				t.Fatalf("Expect error, want %v, got %v", c.expectError, err)
			}
			continue
		}
		if got != c.expect {
			t.Fatalf("Gas estimation mismatch, want %d, got %d", c.expect, got)
		}
	}
}
