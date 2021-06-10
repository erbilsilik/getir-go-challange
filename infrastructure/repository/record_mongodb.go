package repository

import (
	"context"
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordRepository struct {
	Collection *mongo.Collection
}

func NewRecordRepositoryMongoDB() *RecordRepository {
	collection := mongodb.Instance.GetDB().Db.Collection("records")
	return &RecordRepository{
		collection,
	}
}

func (r RecordRepository) List() ([]*entity.Record, error) {
	ctx := context.TODO()
	var records []*entity.Record
	cur, err := r.Collection.Find(ctx, bson.D{{}})
	if err != nil {
		return records, err
	}

	for cur.Next(ctx) {
		var r entity.Record
		err := cur.Decode(&r)
		if err != nil {
			return records, err
		}

		records = append(records, &r)
	}

	if err := cur.Err(); err != nil {
		return records, err
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return records, mongo.ErrNoDocuments
	}

	return records, nil
}