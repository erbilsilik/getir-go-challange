package entity

type Config struct {
	Key        string
	Value 	   interface{}
}

func NewConfig(key string, value interface{}) (*Config, error) {
	c := &Config{
		Key:       key,
		Value:     value,
	}
	err := c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return c, nil
}

func (c *Config) Validate() error {
	if c.Key == "" || c.Value == nil {
		return ErrInvalidEntity
	}
	return nil
}
