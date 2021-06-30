package configuration

import (
	"github.com/erbilsilik/getir-go-challange/entity"
)

//inmem in memory repo
type inmem struct {
	m map[string]*entity.Config
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*entity.Config{}
	return &inmem{
		m: m,
	}
}

func (r *inmem) Create(key string, value interface{}) error{
	config := entity.Config{Key: key, Value: value}
	r.m[config.Key] = &config
	return nil
}

func (r *inmem) FindByKey(key string) *entity.Config {
	return r.m[key]
}