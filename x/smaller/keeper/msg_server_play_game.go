package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tmsdkeys/smaller3/x/smaller/types"
)

func (k msgServer) PlayGame(goCtx context.Context, msg *types.MsgPlayGame) (*types.MsgPlayGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPlayGameResponse{}, nil
}
