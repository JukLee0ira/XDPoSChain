package utils

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/XinFinOrg/XDPoSChain/XDCx"
	"github.com/XinFinOrg/XDPoSChain/XDCxlending"
	"github.com/XinFinOrg/XDPoSChain/eth"
	"github.com/XinFinOrg/XDPoSChain/eth/downloader"
	"github.com/XinFinOrg/XDPoSChain/eth/ethconfig"
	"github.com/XinFinOrg/XDPoSChain/ethstats"
	"github.com/XinFinOrg/XDPoSChain/internal/ethapi"
	"github.com/XinFinOrg/XDPoSChain/metrics"
	"github.com/XinFinOrg/XDPoSChain/node"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *ethconfig.Config, version string) ethapi.Backend {
	if cfg.SyncMode == downloader.LightSync {
		Fatalf("can't register eth service in light sync mode, light mode has been deprecated")
		return nil
	} else {
		// err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// var XDCXServ *XDCx.XDCX
		// var lendingServ *XDCxlending.Lending

		// ctx.Service(&XDCXServ)
		XDCXServ, err := XDCx.New(stack, &XDCx.Config{})
		if err != nil {
			Fatalf("Failed to initialize XDCX service: %v", err)
		}
		// ctx.Service(&lendingServ)
		lendingServ, err := XDCxlending.New(stack, XDCXServ)
		if err != nil {
			Fatalf("Failed to initialize lending service: %v", err)
		}
		backend, err := eth.New(stack, cfg, XDCXServ, lendingServ)
		// 	fullNode, err := eth.New(ctx, cfg, XDCXServ, lendingServ)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		if err != nil {
			Fatalf("Failed to register the Ethereum service: %v", err)
		}
		// TODO: move the following code to function makeFullNode
		// Ref: #21105, #22641, #23761, #24877
		// Create gauge with geth system and build information
		var protos []string
		for _, p := range backend.Protocols() {
			protos = append(protos, fmt.Sprintf("%v/%d", p.Name, p.Version))
		}
		metrics.NewRegisteredGaugeInfo("xdc/info", nil).Update(metrics.GaugeInfoValue{
			"arch":          runtime.GOARCH,
			"os":            runtime.GOOS,
			"version":       version, // cfg.Node.Version
			"eth_protocols": strings.Join(protos, ","),
		})

		// 	return fullNode, err
		// })
		return backend.ApiBackend
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to the node.
func RegisterEthStatsService(stack *node.Node, backend ethapi.Backend, url string) {
	if err := ethstats.New(stack, backend, backend.Engine(), url); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterXDCXService(stack *node.Node, cfg *XDCx.Config) {
	XDCX, err := XDCx.New(stack, cfg)
	// if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
	// 	return XDCX, nil
	// }); err != nil {//TODO:remove this
	if err != nil {
		Fatalf("Failed to register the XDCX service: %v", err)
	}

	// register XDCxlending service//TODO:remove
	// if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
	// 	return XDCxlending.New(XDCX), nil
	// }); err != nil {
	if _, err := XDCxlending.New(stack, XDCX); err != nil {
		Fatalf("Failed to register the XDCXLending service: %v", err)
	}
}
