package repository

import (
	"context"
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/pkg/mongodb"
	"github.com/erbilsilik/getir-go-challange/usecase/record"
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

func (r RecordRepository) GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(
	q *record.RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery,
) ([]*entity.Record, error) {
	ctx := context.TODO()
	var records []*entity.Record

	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{
				{"createdAt", bson.D{
					{"$gte", q.StartDate },
					{"$lte", q.EndDate },
				}},
			}},
		},
		{
			{"$project", bson.D{
				{"key", 1 },
				{"createdAt", 1 },
				{"_id", 0},
				{"totalCount", bson.D{ { "$sum", "$counts"} } },

			}},
		},
		{
			{"$match", bson.D{
				{"totalCount", bson.D{
					{"$lte", q.MaxCount },
					{"$gte", q.MinCount },
				}},
			}},
		},
	}
	cur, err := r.Collection.Aggregate(ctx, pipeline)

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