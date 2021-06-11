package entity

import (
	"time"
)

type Record struct {
	Key        string
	Value 	   string
	CreatedAt  time.Time
	Counts 	   [] int
	TotalCount int
}

func NewRecord(value string, counts []int) (*Record, error) {
	b := &Record{
		Key:       NewID().String(),
		Value:     value,
		Counts:    counts,
		CreatedAt: time.Now(),
	}
	return b, nil
}
