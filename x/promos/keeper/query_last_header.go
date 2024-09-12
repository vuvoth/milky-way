package keeper

import (
	"context"

	"promos/x/promos/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LastHeader(goCtx context.Context, req *types.QueryLastHeaderRequest) (*types.QueryLastHeaderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeApt := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	store := prefix.NewStore(storeApt, types.HeaderPrefix)
	
	heightKey := types.HeaderKey(uint64(req.Height))

	header := store.Get(heightKey)

	return &types.QueryLastHeaderResponse{Header: string(header)}, nil
}
