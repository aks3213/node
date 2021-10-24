package metaclientd

import (
	"fmt"
	"github.com/Meta-Protocol/metacore/app"
	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	stypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	flag "github.com/spf13/pflag"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"strings"
	"sync/atomic"
)

// Broadcast Broadcasts tx to metachain. Returns txHash and error
func (b *MetachainBridge) Broadcast(msgs ...stypes.Msg) (string, error) {
	b.broadcastLock.Lock()
	defer b.broadcastLock.Unlock()
	var err error
	accountNumber, seqNumber, err := b.GetAccountNumberAndSequenceNumber()
	if err != nil {
		return "", err
	}
	b.accountNumber = accountNumber
	if b.seqNumber < seqNumber {
		b.seqNumber = seqNumber
	}
	b.logger.Info().Uint64("account_number", b.accountNumber).Uint64("sequence_number", b.seqNumber).Msg("account info")

	flags := flag.NewFlagSet("metacore", 0)

	ctx := b.GetContext()
	factory := clienttx.NewFactoryCLI(ctx, flags)
	factory = factory.WithAccountNumber(b.accountNumber)
	factory = factory.WithSequence(b.seqNumber)
	factory = factory.WithSignMode(signing.SignMode_SIGN_MODE_DIRECT)

	builder, err := clienttx.BuildUnsignedTx(factory, msgs...)
	if err != nil {
		return "", err
	}
	builder.SetGasLimit(200000000)
	fmt.Printf("signing from name: %s\n", ctx.GetFromName())
	err = clienttx.Sign(factory, ctx.GetFromName(), builder, true)
	if err != nil {
		return "", err
	}

	txBytes, err := ctx.TxConfig.TxEncoder()(builder.GetTx())
	if err != nil {
		return "", err
	}

	// broadcast to a Tendermint node
	commit, err := ctx.BroadcastTxSync(txBytes)
	if err != nil {
		return "", fmt.Errorf("fail to broadcast tx: %w", err)
	}
	// Code will be the tendermint ABICode , it start at 1 , so if it is an error , code will not be zero
	if commit.Code > 0 {
		if commit.Code == 32 {
			// bad sequence number, fetch new one
			_, seqNum, _ := b.GetAccountNumberAndSequenceNumber()
			if seqNum > 0 {
				b.seqNumber = seqNum
			}
		}
		b.logger.Info().Msgf("messages: %+v", msgs)
		return commit.TxHash, fmt.Errorf("fail to broadcast to THORChain,code:%d, log:%s", commit.Code, commit.RawLog)
	}
	b.logger.Info().Msgf("Received a TxHash of %v from the metachain, Code %d, log %s", commit.TxHash, commit.Code, commit.Logs)

	// increment seqNum
	atomic.AddUint64(&b.seqNumber, 1)
	b.logger.Info().Msgf("b.sequence number increased to %d", b.seqNumber)

	return commit.TxHash, nil
}

// GetContext return a valid context with all relevant values set
func (b *MetachainBridge) GetContext() client.Context {
	ctx := client.Context{}
	ctx = ctx.WithKeyring(b.keys.GetKeybase())
	ctx = ctx.WithChainID("metacore")
	ctx = ctx.WithHomeDir(b.cfg.ChainHomeFolder)
	ctx = ctx.WithFromName(b.cfg.SignerName)
	ctx = ctx.WithFromAddress(b.keys.GetSignerInfo().GetAddress())
	ctx = ctx.WithBroadcastMode("sync")

	encodingConfig := app.MakeEncodingConfig()
	ctx = ctx.WithJSONMarshaler(encodingConfig.Marshaler)
	ctx = ctx.WithInterfaceRegistry(encodingConfig.InterfaceRegistry)
	ctx = ctx.WithTxConfig(encodingConfig.TxConfig)
	ctx = ctx.WithLegacyAmino(encodingConfig.Amino)
	ctx = ctx.WithAccountRetriever(authtypes.AccountRetriever{})

	remote := b.cfg.ChainRPC
	if !strings.HasPrefix(b.cfg.ChainHost, "http") {
		remote = fmt.Sprintf("tcp://%s", remote)
	}
	//fmt.Println("ctx.remote ", remote)

	ctx = ctx.WithNodeURI(remote)
	client, err := rpchttp.New(remote, "/websocket")
	if err != nil {
		panic(err)
	}
	ctx = ctx.WithClient(client)
	return ctx
}