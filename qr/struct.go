package qr

type QueryAllNftInfo struct {
	AllNftInfo `json:"all_nft_info"`
}
type AllNftInfo struct {
	TokenId string `json:"token_id"`
}

type ResponseAllNftInfo struct {
}

type QueryResult struct {
}

type Access struct {
	Owner     string
	Approvals []string
}
type Info struct {
}
