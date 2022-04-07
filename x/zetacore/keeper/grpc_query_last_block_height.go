package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zeta-chain/zetacore/x/zetacore/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LastBlockHeightAll(c context.Context, req *types.QueryAllLastBlockHeightRequest) (*types.QueryAllLastBlockHeightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lastBlockHeights []*types.LastBlockHeight
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	lastBlockHeightStore := prefix.NewStore(store, types.KeyPrefix(types.LastBlockHeightKey))

	pageRes, err := query.Paginate(lastBlockHeightStore, req.Pagination, func(key []byte, value []byte) error {
		var lastBlockHeight types.LastBlockHeight
		if err := k.cdc.UnmarshalBinaryBare(value, &lastBlockHeight); err != nil {
			return err
		}

		lastBlockHeights = append(lastBlockHeights, &lastBlockHeight)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLastBlockHeightResponse{LastBlockHeight: lastBlockHeights, Pagination: pageRes}, nil
}

func (k Keeper) LastBlockHeight(c context.Context, req *types.QueryGetLastBlockHeightRequest) (*types.QueryGetLastBlockHeightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLastBlockHeight(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetLastBlockHeightResponse{LastBlockHeight: &val}, nil
}