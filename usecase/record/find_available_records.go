package record

import "github.com/erbilsilik/getir-go-challange/entity"

func List(s *Service) ([]*entity.Record, error) {
	records, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, entity.ErrNotFound
	}
	return records, nil
}
