package types

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
