package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/tmsdkeys/smaller3/testutil/keeper"
	"github.com/tmsdkeys/smaller3/x/smaller/keeper"
	"github.com/tmsdkeys/smaller3/x/smaller/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SmallerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
