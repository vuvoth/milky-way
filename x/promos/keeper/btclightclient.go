package keeper

import (
	"context"
	btclc "promos/x/promos/btclightclient"
	"promos/x/promos/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"cosmossdk.io/store/prefix"
)

// must call in genesis
func (k Keeper) InitLightClient(ctx context.Context, height uint64, lightBlock btclc.BTCLightBlock) error {
	storeApt := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	
	store := prefix.NewStore(storeApt, types.HeaderPrefix)

	key := types.HeaderKey(height)
	store.Set(key, []byte(lightBlock.Data))
	return nil
}

func (k Keeper) UpdateHeader(ctx context.Context, btclc btclc.BTCLightBlock) error {
	
	return nil
}

func (k Keeper) VerifyTransaction(ctx context.Context) error {
	return nil
}
