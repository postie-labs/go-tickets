package ticket

import "github.com/postie-labs/go-crypto-lib"

type Signature struct {
	PubKey *crypto.PubKey `json:"pub_key"`
	Bytes  []byte         `json:"bytes"`
}
