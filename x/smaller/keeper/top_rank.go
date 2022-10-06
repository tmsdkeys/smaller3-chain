package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/tmsdkeys/smaller3/x/smaller/types"
)

// OnRecvTopRankPacket processes packet reception
func (k Keeper) OnRecvTopRankPacket(ctx sdk.Context, packet channeltypes.Packet, data types.TopRankPacketData) (packetAck types.TopRankPacketAck, err error) {
	packetAck = types.TopRankPacketAck{
		ClientId:       "placeholder",
		Address:        data.Address,
		HasBeenTopRank: false,
	}
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic

	return packetAck, nil
}
