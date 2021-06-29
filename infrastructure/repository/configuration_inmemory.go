package repository

import (
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/pkg/caching/inmemory"
	mem "github.com/erbilsilik/getir-go-challange/pkg/caching/inmemory/interfaces"
)

type ConfigRepository struct {
	database mem.InMemoryCachingSource
}

func (c ConfigRepository) Create(key string, value interface{}) error   {
	err := c.database.SetValue(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (c ConfigRepository) FindByKey(key string) *entity.Config {
	config := new(entity.Config)
	data := c.database.GetValueByKey(key)
	if data == nil {
		return config
	}
	config.Key = key
	config.Value = data
	return config
}

func NewConfigurationRepository() *ConfigRepository {
	inmemorySource, _ := inmemory.New()
	return &ConfigRepository{
		database: inmemorySource,
	}
}