package zetaclient

import (
	"context"
	"fmt"
	"sort"

	"time"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/query"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
	"google.golang.org/grpc"
)

type Order string

const (
	NoOrder    Order = ""
	Ascending  Order = "ASC"
	Descending Order = "DESC"
)

func (b *ZetaCoreBridge) GetInboundPermissions() (observertypes.PermissionFlags, error) {
	client := observertypes.NewQueryClient(b.grpcConn)
	resp, err := client.PermissionFlags(context.Background(), &observertypes.QueryGetPermissionFlagsRequest{})
	if err != nil {
		return observertypes.PermissionFlags{}, err
	}
	return resp.PermissionFlags, nil
}

func (b *ZetaCoreBridge) GetCoreParamsForChainID(externalChainID int64) (*observertypes.CoreParams, error) {
	client := observertypes.NewQueryClient(b.grpcConn)
	resp, err := client.GetCoreParamsForChain(context.Background(), &observertypes.QueryGetCoreParamsForChainRequest{ChainId: externalChainID})
	if err != nil {
		return &observertypes.CoreParams{}, err
	}
	return resp.CoreParams, nil
}

func (b *ZetaCoreBridge) GetCoreParams() ([]*observertypes.CoreParams, error) {
	client := observertypes.NewQueryClient(b.grpcConn)
	err := error(nil)
	resp := &observertypes.QueryGetCoreParamsResponse{}
	for i := 0; i <= DefaultRetryCount; i++ {
		resp, err = client.GetCoreParams(context.Background(), &observertypes.QueryGetCoreParamsRequest{})
		if err == nil {
			return resp.CoreParams.CoreParams, nil
		}
		time.Sleep(DefaultRetryInterval * time.Second)
	}
	return nil, fmt.Errorf("failed to get core params | err %s", err.Error())
}

func (b *ZetaCoreBridge) GetObserverParams() (observertypes.Params, error) {
	client := observertypes.NewQueryClient(b.grpcConn)
	resp, err := client.Params(context.Background(), &observertypes.QueryParamsRequest{})
	if err != nil {
		return observertypes.Params{}, err
	}
	return resp.Params, nil
}

func (b *ZetaCoreBridge) GetUpgradePlan() (*upgradetypes.Plan, error) {
	client := upgradetypes.NewQueryClient(b.grpcConn)

	resp, err := client.CurrentPlan(context.Background(), &upgradetypes.QueryCurrentPlanRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Plan, nil
}

//func (b *ZetaCoreBridge) GetAccountDetails(address string) (string, error) {
//	client := authtypes.NewQueryClient(b.grpcConn)
//	resp, err := client.Account(context.Background(), &authtypes.QueryAccountRequest{
//		Address: address,
//	})
//	if err != nil {
//		b.logger.Error().Err(err).Msg("Query account failed")
//		return "", err
//	}
//
//	err := resp.UnpackInterfaces
//	return resp.Account.GetTypeUrl(), nil
//
//}

func (b *ZetaCoreBridge) GetAllCctx() ([]*types.CrossChainTx, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.CctxAll(context.Background(), &types.QueryAllCctxRequest{})
	if err != nil {
		return nil, err
	}
	return resp.CrossChainTx, nil
}

func (b *ZetaCoreBridge) GetCctxByHash(sendHash string) (*types.CrossChainTx, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.Cctx(context.Background(), &types.QueryGetCctxRequest{Index: sendHash})
	if err != nil {
		return nil, err
	}
	return resp.CrossChainTx, nil
}

func (b *ZetaCoreBridge) GetCctxByNonce(chainID int64, nonce uint64) (*types.CrossChainTx, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.CctxByNonce(context.Background(), &types.QueryGetCctxByNonceRequest{
		ChainID: chainID,
		Nonce:   nonce,
	})
	if err != nil {
		return nil, err
	}
	return resp.CrossChainTx, nil
}

func (b *ZetaCoreBridge) GetObserverList(chain common.Chain) ([]string, error) {
	client := observertypes.NewQueryClient(b.grpcConn)
	resp, err := client.ObserversByChain(context.Background(), &observertypes.QueryObserversByChainRequest{
		ObservationChain: chain.ChainName.String(),
	})
	if err != nil {
		return nil, err
	}
	return resp.Observers, nil
}

func (b *ZetaCoreBridge) GetAllPendingCctx(chainID uint64) ([]*types.CrossChainTx, error) {
	client := types.NewQueryClient(b.grpcConn)
	maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 1024 * 1024)
	resp, err := client.CctxAllPending(context.Background(), &types.QueryAllCctxPendingRequest{ChainId: chainID}, maxSizeOption)
	if err != nil {
		return nil, err
	}
	return resp.CrossChainTx, nil
}

func (b *ZetaCoreBridge) GetLastBlockHeight() ([]*types.LastBlockHeight, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.LastBlockHeightAll(context.Background(), &types.QueryAllLastBlockHeightRequest{})
	if err != nil {
		b.logger.Error().Err(err).Msg("query GetBlockHeight error")
		return nil, err
	}
	return resp.LastBlockHeight, nil
}

func (b *ZetaCoreBridge) GetLatestZetaBlock() (*tmtypes.Block, error) {
	client := tmservice.NewServiceClient(b.grpcConn)
	res, err := client.GetLatestBlock(context.Background(), &tmservice.GetLatestBlockRequest{})
	if err != nil {
		return nil, err
	}
	return res.Block, nil
}

func (b *ZetaCoreBridge) GetLastBlockHeightByChain(chain common.Chain) (*types.LastBlockHeight, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.LastBlockHeight(context.Background(), &types.QueryGetLastBlockHeightRequest{Index: chain.ChainName.String()})
	if err != nil {
		return nil, err
	}
	return resp.LastBlockHeight, nil
}

func (b *ZetaCoreBridge) GetZetaBlockHeight() (int64, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.LastZetaHeight(context.Background(), &types.QueryLastZetaHeightRequest{})
	if err != nil {
		return 0, err
	}
	return resp.Height, nil
}

func (b *ZetaCoreBridge) GetNonceByChain(chain common.Chain) (*types.ChainNonces, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.ChainNonces(context.Background(), &types.QueryGetChainNoncesRequest{Index: chain.ChainName.String()})
	if err != nil {
		return nil, err
	}
	return resp.ChainNonces, nil
}

func (b *ZetaCoreBridge) GetAllNodeAccounts() ([]*observertypes.NodeAccount, error) {
	client := observertypes.NewQueryClient(b.grpcConn)
	resp, err := client.NodeAccountAll(context.Background(), &observertypes.QueryAllNodeAccountRequest{})
	if err != nil {
		return nil, err
	}
	b.logger.Debug().Msgf("GetAllNodeAccounts: %d", len(resp.NodeAccount))
	return resp.NodeAccount, nil
}

func (b *ZetaCoreBridge) GetKeyGen() (*observertypes.Keygen, error) {
	client := observertypes.NewQueryClient(b.grpcConn)
	resp, err := client.Keygen(context.Background(), &observertypes.QueryGetKeygenRequest{})
	if err != nil {
		//log.Error().Err(err).Msg("query GetKeyGen error")
		return nil, err
	}
	return resp.Keygen, nil
}

func (b *ZetaCoreBridge) GetOutTxTracker(chain common.Chain, nonce uint64) (*types.OutTxTracker, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.OutTxTracker(context.Background(), &types.QueryGetOutTxTrackerRequest{
		ChainID: chain.ChainId,
		Nonce:   nonce,
	})
	if err != nil {
		return nil, err
	}
	return &resp.OutTxTracker, nil
}

func (b *ZetaCoreBridge) GetAllOutTxTrackerByChain(chain common.Chain, order Order) ([]types.OutTxTracker, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.OutTxTrackerAllByChain(context.Background(), &types.QueryAllOutTxTrackerByChainRequest{
		Chain: chain.ChainId,
		Pagination: &query.PageRequest{
			Key:        nil,
			Offset:     0,
			Limit:      300,
			CountTotal: false,
			Reverse:    false,
		},
	})
	if err != nil {
		return nil, err
	}
	if order == Ascending {
		sort.SliceStable(resp.OutTxTracker, func(i, j int) bool {
			return resp.OutTxTracker[i].Nonce < resp.OutTxTracker[j].Nonce
		})
	}
	if order == Descending {
		sort.SliceStable(resp.OutTxTracker, func(i, j int) bool {
			return resp.OutTxTracker[i].Nonce > resp.OutTxTracker[j].Nonce
		})
	}
	return resp.OutTxTracker, nil
}

func (b *ZetaCoreBridge) GetClientParams(chainID int64) (observertypes.QueryGetCoreParamsForChainResponse, error) {
	client := observertypes.NewQueryClient(b.grpcConn)
	resp, err := client.GetCoreParamsForChain(context.Background(), &observertypes.QueryGetCoreParamsForChainRequest{ChainId: chainID})
	if err != nil {
		return observertypes.QueryGetCoreParamsForChainResponse{}, err
	}
	return *resp, nil
}

func (b *ZetaCoreBridge) GetPendingNonces() (*types.QueryAllPendingNoncesResponse, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.PendingNoncesAll(context.Background(), &types.QueryAllPendingNoncesRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
