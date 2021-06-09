package entity

import (
	"time"
)

type Record struct {
	Key        string
	TotalCount int
	CreatedAt  time.Time
}

func NewRecord(key string, totalCount int) (*Record, error) {
	r := &Record{
		Key:        key,
		TotalCount: totalCount,
		CreatedAt:  time.Now(),
	}
	return r, nil
}
