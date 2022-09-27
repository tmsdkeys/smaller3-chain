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

func createNStoredGame(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredGame {
	items := make([]types.StoredGame, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredGame(ctx, items[i])
	}
	return items
}

func TestStoredGameGet(t *testing.T) {
	keeper, ctx := keepertest.SmallerKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredGame(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredGameRemove(t *testing.T) {
	keeper, ctx := keepertest.SmallerKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredGame(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredGame(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredGameGetAll(t *testing.T) {
	keeper, ctx := keepertest.SmallerKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredGame(ctx)),
	)
}