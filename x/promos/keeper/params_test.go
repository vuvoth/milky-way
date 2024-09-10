package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "promos/testutil/keeper"
	"promos/x/promos/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PromosKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
