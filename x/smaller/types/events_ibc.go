package types

// IBC events
const (
	EventTypeTimeout          = "timeout"
	EventTypeGameResultPacket = "gameResult_packet"
	EventTypeTopRankPacket    = "topRank_packet"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
