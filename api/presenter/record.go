package presenter

type Record struct {
	Key        string `json:"key"`
	TotalCount int    `json:"totalCount"`
	CreatedAt  string `json:"createdAt"`
}