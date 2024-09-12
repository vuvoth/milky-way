package simulation

import (
	"math/rand"

	"promos/x/promos/keeper"
	"promos/x/promos/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgInitLightClient(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgInitLightClient{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the InitLightClient simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "InitLightClient simulation not implemented"), nil, nil
	}
}
