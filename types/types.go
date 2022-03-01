package types

import (
	"encoding/hex"
	"time"
)

var (
	EmptyHash = Hash{}
)

type (
	Hash      [32]byte
	Timestamp = time.Time
	Data      []byte
)

func (hash *Hash) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hex.EncodeToString(hash[:]) + `"`), nil
}

func (hash *Hash) UnmarshalJSON(data []byte) error {
	tmp, err := hex.DecodeString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	copy(hash[:], tmp)
	return nil
}

func TimestampNow() Timestamp {
	return time.Now().Round(0) // remove milliseconds
}
