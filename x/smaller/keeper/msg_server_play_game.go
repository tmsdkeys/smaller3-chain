package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tmsdkeys/smaller3/x/smaller/types"
)

func (k msgServer) PlayGame(goCtx context.Context, msg *types.MsgPlayGame) (*types.MsgPlayGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found, cannot continue")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	numPick, err := strconv.ParseUint(msg.NumPick, 10, 64)
	if err != nil {
		// TODO: change to correct error
		return nil, sdkerrors.Wrapf(types.ErrInvalidLengthStoredGame, msg.NumPick)
	}

	win := numPick < 3

	storedGame := types.StoredGame{
		Index:    newIndex,
		Player:   msg.Creator,
		Win:      win,
		SentToLB: false,
	}

	k.Keeper.SetStoredGame(ctx, storedGame)

	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	return &types.MsgPlayGameResponse{
		GameIndex: newIndex,
		Win:       win,
	}, nil
}
