package record

import "github.com/erbilsilik/getir-go-challange/entity"

func (s *Service) CreateRecord(key string, totalCount int) error {
	if r, err := entity.NewRecord(key, totalCount); err != nil {
		return err
	} else {
		return s.repo.Create(r)
	}
}
