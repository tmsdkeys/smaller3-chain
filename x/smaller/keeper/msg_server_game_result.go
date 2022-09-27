package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/tmsdkeys/smaller3/x/smaller/types"
)

func (k msgServer) SendGameResult(goCtx context.Context, msg *types.MsgSendGameResult) (*types.MsgSendGameResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.GameResultPacketData

	packet.GameId = msg.GameId

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
