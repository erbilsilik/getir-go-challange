package configuration

import (
	"github.com/erbilsilik/getir-go-challange/entity"
)

type Reader interface {
	FindByKey(query string) *entity.Config
}

type Writer interface {
	Create(key string, value interface{}) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Create(key string, value interface{}) error
	FindByKey(query string) *entity.Config
}
