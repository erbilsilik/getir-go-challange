package record

import "github.com/erbilsilik/getir-go-challange/entity"

func (s *Service) FindAvailableRecords(q *FindAvailableRecordsQuery) ([]*entity.RecordTotalCount, error) {
	records, err := s.repo.List(q)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, entity.ErrNotFound
	}
	return records, nil
}
