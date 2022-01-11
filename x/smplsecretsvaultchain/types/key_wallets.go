package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WalletsKeyPrefix is the prefix to retrieve all Wallets
	WalletsKeyPrefix = "Wallets/value/"
)

// WalletsKey returns the store key to retrieve a Wallets from the index fields
func WalletsKey(
	owner string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
