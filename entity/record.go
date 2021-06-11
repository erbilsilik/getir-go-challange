package entity

import "time"

type Record struct {
	Key        string
	Value 	   string
	CreatedAt  time.Time
	Counts 	   [] int
}

type RecordTotalCount struct {
	Key string
	TotalCount int
	CreatedAt  time.Time
}
