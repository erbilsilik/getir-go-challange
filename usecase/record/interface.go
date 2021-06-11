package record

import (
	"github.com/erbilsilik/getir-go-challange/entity"
	"time"
)

type Reader interface {
	List(query *FindAvailableRecordsQuery) ([]*entity.RecordTotalCount, error)
}

type Writer interface {}

type Repository interface {
	Reader
	Writer
}

type FindAvailableRecordsQuery struct {
	StartDate string
	EndDate   string
	MinCount  int
	MaxCount  int
}

type RecordTotalCount struct {
	Key string
	TotalCount int
	CreatedAt  time.Time
}

type UseCase interface {
	List(query *FindAvailableRecordsQuery) ([]*entity.RecordTotalCount, error)
}
