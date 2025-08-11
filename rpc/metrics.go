// Copyright 2020 The go-ethereum Authors
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

package rpc

import (
	"fmt"
	"reflect"
	"time"

	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/common/hexutil"
	"github.com/XinFinOrg/XDPoSChain/log"
	"github.com/XinFinOrg/XDPoSChain/metrics"
)

var (
	rpcRequestGauge        = metrics.NewRegisteredGauge("rpc/requests", nil)
	successfulRequestGauge = metrics.NewRegisteredGauge("rpc/success", nil)
	failedRequestGauge     = metrics.NewRegisteredGauge("rpc/failure", nil)

	// serveTimeHistName is the prefix of the per-request serving time histograms.
	serveTimeHistName = "rpc/duration"

	rpcServingTimer = metrics.NewRegisteredTimer("rpc/duration/all", nil)
)

// updateServeTimeHistogram tracks the serving time of a remote RPC call.
func updateServeTimeHistogram(method string, success bool, elapsed time.Duration, params ...interface{}) {
	note := "success"
	if !success {
		note = "failure"
	}
	h := fmt.Sprintf("%s/%s/%s", serveTimeHistName, method, note)
	sampler := func() metrics.Sample {
		return metrics.ResettingSample(
			metrics.NewExpDecaySample(1028, 0.015),
		)
	}
	metrics.GetOrRegisterHistogramLazy(h, nil, sampler).Update(elapsed.Nanoseconds())

	// Add metrics for eth_call with contract/caller info
	if method == "eth_call" && len(params) > 0 {
		log.Debug("eth_call paramsvvvvvvvvvvvv", "params", params)
		log.Debug("eth_call params[0] type", "type", fmt.Sprintf("%T", params[0]))

		// Use reflection to access the From and To fields
		v := reflect.ValueOf(params[0])
		if v.Kind() == reflect.Struct || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct) {
			// If it's a pointer, get the struct it points to
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}

			baseMetric := fmt.Sprintf("%s/eth_call", h)

			// Get To field
			// rpc/duration/eth_call/success.xdc0000000000000000000000000000000000000088
			if toField := v.FieldByName("To"); toField.IsValid() && !toField.IsNil() {
				to := toField.Interface().(*common.Address)
				log.Debug("eth_call contract addressvvvvvvvvvvvv", "to", to.Hex())
				// Record contract address calls
				contractMetric := fmt.Sprintf("%s/contract/%s", baseMetric, to.Hex())
				metrics.GetOrRegisterMeter(contractMetric, nil).Mark(1)

				// Get Data field for function signature tracking
				var funcSig string
				if dataField := v.FieldByName("Data"); dataField.IsValid() && !dataField.IsNil() {
					data := dataField.Interface().(*hexutil.Bytes)
					if data != nil && len(*data) >= 4 {
						// Extract first 4 bytes (8 hex characters) as function signature
						funcSig = hexutil.Encode((*data)[:4])
						log.Debug("eth_call function signature", "func_sig", funcSig)
					}
				} else if inputField := v.FieldByName("Input"); inputField.IsValid() && !inputField.IsNil() {
					input := inputField.Interface().(*hexutil.Bytes)
					if input != nil && len(*input) >= 4 {
						// Extract first 4 bytes (8 hex characters) as function signature
						funcSig = hexutil.Encode((*input)[:4])
						log.Debug("eth_call function signature", "func_sig", funcSig)
					}
				}

				// Create detailed metric with both contract address and function signature
				if funcSig != "" {
					detailedMetric := fmt.Sprintf("%s/contract/%s/func/%s", baseMetric, to.Hex(), funcSig)
					metrics.GetOrRegisterMeter(detailedMetric, nil).Mark(1)
				}
			}

		} else {
			log.Debug("eth_call paramsvvvvvvvvvvvv,no ok!", "params", params)
		}
	}
}
