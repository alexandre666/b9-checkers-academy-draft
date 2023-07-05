package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredenergyKeyPrefix is the prefix to retrieve all Storedenergy
	StoredenergyKeyPrefix = "Storedenergy/value/"
)

// StoredenergyKey returns the store key to retrieve a Storedenergy from the index fields
func StoredenergyKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
