package msgs

import "github.com/postie-labs/go-tickets/types"

type QueryAllNftInfo struct {
	AllNftInfo `json:"all_nft_info"`
}
type AllNftInfo struct {
	TokenId string `json:"token_id"`
}
type ResponseAllNftInfo struct {
	QueryResult `json:"query_result"`
}
type QueryResult struct {
	types.Access `json:"access"`
	types.Info   `json:"info"`
}
