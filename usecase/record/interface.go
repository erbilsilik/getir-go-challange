package record

import (
	"github.com/erbilsilik/getir-go-challange/entity"
	"time"
)

type Reader interface {
	GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(
		query *RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery,
	) ([]*entity.Record, error)
}

type Writer interface {}

type Repository interface {
	Reader
	Writer
}

type RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery struct {
	StartDate time.Time
	EndDate   time.Time
	MinCount  int
	MaxCount  int
}

type UseCase interface {
	GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(
		query *RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery,
	) ([]*entity.Record, error)
}
