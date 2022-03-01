package blockchain

import (
	"github.com/postie-labs/go-postie-lib/crypto"
)

type Signature struct {
	PubKey *crypto.PubKey `json:"pub_key"`
	Bytes  []byte         `json:"bytes"`
}
