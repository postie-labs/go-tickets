package types

import (
	"bytes"
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

func NewHashFromString(str string) (Hash, error) {
	var hash Hash
	tmp, err := hex.DecodeString(str)
	if err != nil {
		return hash, err
	}
	copy(hash[:], tmp)
	return hash, nil
}

func (hash *Hash) String() string {
	return hex.EncodeToString(hash[:])
}

func (hash *Hash) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hash.String() + `"`), nil
}

func (hash *Hash) UnmarshalJSON(data []byte) error {
	tmp, err := NewHashFromString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	*hash = tmp
	return nil
}

func (hash *Hash) IsEmpty() bool {
	return hash.Equals(EmptyHash)
}

func (hash *Hash) Equals(target Hash) bool {
	return bytes.Equal(hash[:], target[:])
}

func TimestampNow() Timestamp {
	return time.Now().Round(0) // remove milliseconds
}
