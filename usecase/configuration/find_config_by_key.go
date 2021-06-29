package configuration

import "github.com/erbilsilik/getir-go-challange/entity"

func (s *Service) FindByKey(key string) *entity.Config {
	config := s.repo.FindByKey(key)
	return config
}
