package record

import (
	"github.com/erbilsilik/getir-go-challange/entity"
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

type UseCase interface {
	List(query *FindAvailableRecordsQuery) ([]*entity.RecordTotalCount, error)
}
