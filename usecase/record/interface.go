package record

import (
	"github.com/erbilsilik/getir-go-challange/entity"
	"time"
)

type Reader interface {
	CalculateRecordsTotalCount(query *CalculateRecordsTotalCountQuery) ([]*entity.Record, error)
}

type Writer interface {}

type Repository interface {
	Reader
	Writer
}

type CalculateRecordsTotalCountQuery struct {
	StartDate time.Time
	EndDate   time.Time
	MinCount  int
	MaxCount  int
}

type UseCase interface {
	CalculateRecordsTotalCount(query *CalculateRecordsTotalCountQuery) ([]*entity.Record, error)
}
