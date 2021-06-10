package record

import "github.com/erbilsilik/getir-go-challange/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) List() ([]*entity.Record, error) {
	records, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, entity.ErrNotFound
	}
	return records, nil
}
