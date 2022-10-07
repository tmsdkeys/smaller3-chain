package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tmsdkeys/smaller3/testutil/keeper"
	"github.com/tmsdkeys/smaller3/testutil/nullify"
	"github.com/tmsdkeys/smaller3/x/smaller/keeper"
	"github.com/tmsdkeys/smaller3/x/smaller/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPlayerInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PlayerInfo {
	items := make([]types.PlayerInfo, n)
	for i := range items {
		items[i].Player = strconv.Itoa(i)

		keeper.SetPlayerInfo(ctx, items[i])
	}
	return items
}

func TestPlayerInfoGet(t *testing.T) {
	keeper, ctx := keepertest.SmallerKeeper(t)
	items := createNPlayerInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPlayerInfo(ctx,
			item.Player,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPlayerInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.SmallerKeeper(t)
	items := createNPlayerInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePlayerInfo(ctx,
			item.Player,
		)
		_, found := keeper.GetPlayerInfo(ctx,
			item.Player,
		)
		require.False(t, found)
	}
}

func TestPlayerInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.SmallerKeeper(t)
	items := createNPlayerInfo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPlayerInfo(ctx)),
	)
}
