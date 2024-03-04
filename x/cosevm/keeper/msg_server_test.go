package keeper_test

import (
	"context"
	"testing"

	keepertest "cosevm/testutil/keeper"
	"cosevm/x/cosevm/keeper"
	"cosevm/x/cosevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosevmKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
