package types

// ValidateBasic is used for validating the packet
func (p TopRankPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p TopRankPacketData) GetBytes() ([]byte, error) {
	var modulePacket SmallerPacketData

	modulePacket.Packet = &SmallerPacketData_TopRankPacket{&p}

	return modulePacket.Marshal()
}
