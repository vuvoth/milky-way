package keeper

import "context"
import 	(btclc "promos/x/promos/btclightclient")

// must call in genesis
func (keeper Keeper) InitLightClient(ctx context.Context, lightBlock btclc.BTCLightBlock) error {
	return nil
}

func (keeper Keeper) UpdateHeader(ctx context.Context, btclcbtclc.BTCLightBlock) error {
	return nil
}

func (keeper Keeper) VerifyTransaction(ctx context.Context) error {
	return nil
}
