package record

import "github.com/erbilsilik/getir-go-challange/entity"

func (s *Service) GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(
	q *RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery) ([]*entity.Record, error,
) {
	records, err := s.repo.GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(q)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, entity.ErrNotFound
	}
	return records, nil
}
