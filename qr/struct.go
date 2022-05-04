package qr

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
	Access `json:"access"`
	Info   `json:"info"`
}

type Access struct {
	Owner     string   `json:"owner"`
	Approvals []string `json:"approvals"`
}
type Info struct {
	TokenUri  string `json:"token_uri"`
	Extension `json:"extension"`
}
type Extension struct {
	NotValidBefore int64       `json:"not_valid_before"`
	NotValidAfter  int64       `json:"not_valid_after"`
	Attributes     []Attribute `json:"attributes"`
}
type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TxExecuteMint struct {
	Mint `json:"mint"`
}
type Mint struct {
	Owner     string    `json:"owner"`
	TokenId   string    `json:"token_id"`
	TokenUri  string    `json:"token_uri"`
	Extension Extension `json:"extension"`
}
