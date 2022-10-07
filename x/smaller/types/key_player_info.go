package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PlayerInfoKeyPrefix is the prefix to retrieve all PlayerInfo
	PlayerInfoKeyPrefix = "PlayerInfo/value/"
)

// PlayerInfoKey returns the store key to retrieve a PlayerInfo from the index fields
func PlayerInfoKey(
	player string,
) []byte {
	var key []byte

	playerBytes := []byte(player)
	key = append(key, playerBytes...)
	key = append(key, []byte("/")...)

	return key
}
