package record

import (
	"github.com/erbilsilik/getir-go-challange/entity"
)

type Reader interface {
	CalculateRecordsTotalCount(query *CalculateRecordsTotalCountQuery) ([]*entity.RecordTotalCount, error)
}

type Writer interface {}

type Repository interface {
	Reader
	Writer
}

type CalculateRecordsTotalCountQuery struct {
	StartDate string
	EndDate   string
	MinCount  int
	MaxCount  int
}

type UseCase interface {
	CalculateRecordsTotalCount(query *CalculateRecordsTotalCountQuery) ([]*entity.RecordTotalCount, error)
}
