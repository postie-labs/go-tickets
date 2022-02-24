package util

import (
	"crypto/sha256"

	"github.com/postie-labs/go-tickets/types"
)

func ToSHA256(data []byte) types.Hash {
	return sha256.Sum256(data)
}
