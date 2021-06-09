package mongodb

import (
	"context"
	"fmt"
	"github.com/erbilsilik/getir-go-challange/api/presenter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongodbDatabase struct {
	db *mongo.Collection
}

func (m mongodbDatabase) Insert(ctx context.Context, entity interface{}) error {
	if _, err := m.db.InsertOne(ctx, entity); err != nil {
		return err
	} else {
		return nil
	}
}

func (m mongodbDatabase) FindAll(ctx context.Context) {
	cur, err := m.db.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		var r presenter.Record
		err := cur.Decode(&r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(r)
	}
}

//func (m mongodbDatabase) FindByFilter(ctx context.Context, entity interface{}, condition string, params ...interface{}) error {
//	if cursor, err := m.db.Find(ctx, bson.M{condition: params}); err != nil {
//		return err
//	} else {
//		var records []interface{}
//		defer cursor.Close(ctx)
//		for cursor.Next(ctx) {
//			var record interface{}
//			if err = cursor.Decode(&record); err != nil {
//				// ignored
//			}
//			records = append(records, record)
//		}
//		entity = records
//	}
//}

func New(connectionString string, database string, collection string) (*mongodbDatabase, error) {
	if client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString)); err != nil {
		return nil, err
	} else {
		return &mongodbDatabase{
			db: client.Database(database).Collection(collection),
		}, nil
	}
}
