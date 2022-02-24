package types

import "time"

type (
	Hash      [32]byte
	Timestamp = time.Time
	Data      []byte
)

func TimestampNow() Timestamp {
	return time.Now()
}
