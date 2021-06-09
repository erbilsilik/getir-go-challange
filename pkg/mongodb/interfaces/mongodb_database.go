package interfaces

import "context"

type MongoDbDatabase interface {
	Insert(ctx context.Context, entity interface{}) error
	FindByFilter(ctx context.Context, entity interface{}, query string, params ...interface{}) error
}
