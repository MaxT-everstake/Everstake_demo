package keeper_test

import (
	"context"
	"testing"

	keepertest "Everstake_demo/testutil/keeper"
	"Everstake_demo/x/everstakedemo/keeper"
	"Everstake_demo/x/everstakedemo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EverstakedemoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
