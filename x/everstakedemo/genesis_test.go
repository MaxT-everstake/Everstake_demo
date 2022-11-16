package everstakedemo_test

import (
	"testing"

	keepertest "Everstake_demo/testutil/keeper"
	"Everstake_demo/testutil/nullify"
	"Everstake_demo/x/everstakedemo"
	"Everstake_demo/x/everstakedemo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EverstakedemoKeeper(t)
	everstakedemo.InitGenesis(ctx, *k, genesisState)
	got := everstakedemo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
