package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/tmsdkeys/smaller3/x/smaller/types"
)

func (k msgServer) SendGameResult(goCtx context.Context, msg *types.MsgSendGameResult) (*types.MsgSendGameResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet
	storedGame, found := k.GetStoredGame(ctx, strconv.FormatUint(msg.GameId, 10))
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Could not find game")
	}

	playerInfo, found := k.GetPlayerInfo(ctx, storedGame.Player)
	// Construct the packet
	var packet types.GameResultPacketData

	packet.GameId = msg.GameId
	packet.Address = storedGame.Player
	packet.WinCount = playerInfo.WinCount

	// Transmit the packet
	err := k.TransmitGameResultPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendGameResultResponse{}, nil
}
