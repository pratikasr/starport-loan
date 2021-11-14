package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pratikasr/starport-loan/testutil/keeper"
	"github.com/pratikasr/starport-loan/x/starportloan/keeper"
	"github.com/pratikasr/starport-loan/x/starportloan/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.StarportloanKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
