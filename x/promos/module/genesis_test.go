package promos_test

import (
	"testing"

	keepertest "promos/testutil/keeper"
	"promos/testutil/nullify"
	promos "promos/x/promos/module"
	"promos/x/promos/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PromosKeeper(t)
	promos.InitGenesis(ctx, k, genesisState)
	got := promos.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
