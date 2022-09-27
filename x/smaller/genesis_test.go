package smaller_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tmsdkeys/smaller3/testutil/keeper"
	"github.com/tmsdkeys/smaller3/testutil/nullify"
	"github.com/tmsdkeys/smaller3/x/smaller"
	"github.com/tmsdkeys/smaller3/x/smaller/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		SystemInfo: types.SystemInfo{
			NextId: 20,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SmallerKeeper(t)
	smaller.InitGenesis(ctx, *k, genesisState)
	got := smaller.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	// this line is used by starport scaffolding # genesis/test/assert
}
