package repository

import (
	"context"
	"fmt"
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/pkg/mongodb"
	"github.com/erbilsilik/getir-go-challange/usecase/record"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// TODO move this util method
func parseDate(layout string, dateString string) time.Time {
	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		panic("parsing date error")
	}
	return parsedDate
}

type RecordRepository struct {
	Collection *mongo.Collection
}

func NewRecordRepositoryMongoDB() *RecordRepository {
	collection := mongodb.Instance.GetDB().Db.Collection("records")
	return &RecordRepository{
		collection,
	}
}

func (r RecordRepository) List(q *record.FindAvailableRecordsQuery) ([]*entity.RecordTotalCount, error) {
	ctx := context.TODO()
	var records []*entity.RecordTotalCount
	//findOptions := options.Find()
	//findOptions.SetLimit(20)
	//findOptions.SetSkip(2)

	layout := "2006-01-02"
	startDateParsed := parseDate(layout, q.StartDate)
	endDateParsed := parseDate(layout, q.EndDate)

	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{
				{"createdAt", bson.D{
					{"$gte", startDateParsed },
					{"$lte", endDateParsed },
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
		fmt.Println(err)
		//return records, err
	}

	for cur.Next(ctx) {
		var r entity.RecordTotalCount
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