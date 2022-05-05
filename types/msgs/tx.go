package msgs

import "github.com/postie-labs/go-tickets/types"

type TxExecuteMint struct {
	Mint `json:"mint"`
}
type Mint struct {
	Owner     string          `json:"owner"`
	TokenId   string          `json:"token_id"`
	TokenUri  string          `json:"token_uri"`
	Extension types.Extension `json:"extension"`
}
