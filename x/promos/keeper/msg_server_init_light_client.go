package keeper

import (
	"context"

	"promos/x/promos/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InitLightClient(goCtx context.Context, msg *types.MsgInitLightClient) (*types.MsgInitLightClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storeApt := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx));
	store := prefix.NewStore(storeApt, types.HeaderPrefix);

	heightKey := types.HeaderKey(msg.Height);

	store.Set(heightKey, []byte(msg.LightBlock));
	return &types.MsgInitLightClientResponse{}, nil
}
