package types

import "time"

type (
	Hash      [32]byte
	Timestamp time.Time
	Data      []byte

	Address   [20]byte // TODO: design concrete Address
	Signature []byte   // TODO: design concrete crypto.Signature
)
