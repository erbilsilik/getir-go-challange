package record

import "github.com/erbilsilik/getir-go-challange/entity"

func (s *Service) CalculateRecordsTotalCount(q *CalculateRecordsTotalCountQuery) ([]*entity.RecordTotalCount, error) {
	records, err := s.repo.CalculateRecordsTotalCount(q)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, entity.ErrNotFound
	}
	return records, nil
}