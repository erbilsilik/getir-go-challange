package configuration

func (s *Service) Create(key string, value interface{}) error {
	err := s.repo.Create(key, value)
	if err != nil {
		return err
	}
	return nil
}
