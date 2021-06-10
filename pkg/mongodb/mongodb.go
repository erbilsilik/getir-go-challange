package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbDatabase struct {
	Db *mongo.Database
}

var Instance *mongodbDatabase

func (mDB *mongodbDatabase) GetDB() *mongodbDatabase {
	return Instance
}

func New(connectionString string, database string) {
	if Instance != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString)); err != nil {
		fmt.Println(err)
	} else {
		Instance = &mongodbDatabase{
			Db: client.Database(database),
		}
	}
}
