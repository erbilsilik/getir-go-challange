package record

import "github.com/erbilsilik/getir-go-challange/entity"

type Reader interface {
	List() ([]*entity.Record, error)
}

type Writer interface {
	Create(e *entity.Record) error
}

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
	ListRecords(query *FindAvailableRecordsQuery) ([]*entity.Record, error)
	CreateRecord(key string, totalCount int) error
}
