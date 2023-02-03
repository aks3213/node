package main

import (
	"context"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/zeta-chain/zetacore/contracts/evm/zetaconnectoreth"
	"math/big"
	"time"
)

func (sm *SmokeTest) TestMessagePassing() {
	startTime := time.Now()
	defer func() {
		fmt.Printf("test finishes in %s\n", time.Since(startTime))
	}()
	// ==================== Interacting with contracts ====================
	time.Sleep(10 * time.Second)
	LoudPrintf("Goerli->Goerli Message Passing (Sending ZETA only)\n")
	fmt.Printf("Approving ConnectorEth to spend deployer's ZetaEth\n")
	amount := big.NewInt(1e18)
	amount = amount.Mul(amount, big.NewInt(10)) // 10 Zeta
	auth := sm.goerliAuth
	tx, err := sm.ZetaEth.Approve(auth, sm.ConnectorEthAddr, amount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Approve tx hash: %s\n", tx.Hash().Hex())
	receipt := MustWaitForTxReceipt(sm.goerliClient, tx)
	fmt.Printf("Approve tx receipt: %d\n", receipt.Status)
	fmt.Printf("Calling ConnectorEth.Send\n")
	tx, err = sm.ConnectorEth.Send(auth, zetaconnectoreth.ZetaInterfacesSendInput{
		DestinationChainId:  big.NewInt(1337), // in dev mode, GOERLI has chainid 1337
		DestinationAddress:  DeployerAddress.Bytes(),
		DestinationGasLimit: big.NewInt(250_000),
		Message:             nil,
		ZetaValueAndGas:     amount,
		ZetaParams:          nil,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ConnectorEth.Send tx hash: %s\n", tx.Hash().Hex())
	receipt = MustWaitForTxReceipt(sm.goerliClient, tx)
	fmt.Printf("ConnectorEth.Send tx receipt: status %d\n", receipt.Status)
	fmt.Printf("  Logs:\n")
	for _, log := range receipt.Logs {
		sentLog, err := sm.ConnectorEth.ParseZetaSent(*log)
		if err == nil {
			fmt.Printf("    Dest Addr: %s\n", ethcommon.BytesToAddress(sentLog.DestinationAddress).Hex())
			fmt.Printf("    Dest Chain: %d\n", sentLog.DestinationChainId)
			fmt.Printf("    Dest Gas: %d\n", sentLog.DestinationGasLimit)
			fmt.Printf("    Zeta Value: %d\n", sentLog.ZetaValueAndGas)
		}
	}
	sm.wg.Add(1)
	go func() {
		defer sm.wg.Done()
		fmt.Printf("Waiting for ConnectorEth.Send CCTX to be mined...\n")
		cctx := WaitCctxMinedByInTxHash(receipt.TxHash.String(), sm.cctxClient)
		receipt, err := sm.goerliClient.TransactionReceipt(context.Background(), ethcommon.HexToHash(cctx.OutboundTxParams.OutboundTxHash))
		if err != nil {
			panic(err)
		}
		for _, log := range receipt.Logs {
			event, err := sm.ConnectorEth.ParseZetaReceived(*log)
			if err == nil {
				fmt.Printf("Received ZetaSent event:\n")
				fmt.Printf("  Dest Addr: %s\n", event.DestinationAddress)
				fmt.Printf("  Zeta Value: %d\n", event.ZetaValue)
				fmt.Printf("  src chainid: %d\n", event.SourceChainId)
				if event.ZetaValue.Cmp(cctx.ZetaMint.BigInt()) != 0 {
					panic("Zeta value mismatch")
				}
			}
		}
	}()
	sm.wg.Wait()
}