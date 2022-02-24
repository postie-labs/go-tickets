package ticket

import "github.com/postie-labs/go-crypto-lib"

type Signature struct {
	PubKey   *crypto.PubKey `json:"pubkey"`
	SigBytes []byte         `json:"sig_bytes"`
}
