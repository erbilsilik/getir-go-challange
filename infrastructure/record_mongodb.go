package repository

import (
	"context"

	"github.com/erbilsilik/getir-go-challange/entity"
	mongo "github.com/erbilsilik/getir-go-challange/pkg/mongodb/interfaces"
)

var ctx = context.TODO()

type recordRepository struct {
	database mongo.MongoDbDatabase
}

func (r recordRepository) Create(record *entity.Record) {
	err := r.database.Insert(ctx, record)
	if err != nil {
		return
	}
}
