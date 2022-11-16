package keeper_test

import (
	"testing"

	testkeeper "Everstake_demo/testutil/keeper"
	"Everstake_demo/x/everstakedemo/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EverstakedemoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
