package presenter

import "time"

type Record struct {
	Key        string `json:"key"`
	TotalCount int    `json:"totalCount"`
	CreatedAt  time.Time `json:"createdAt"`
}