package interfaces

type MongoDbDatabase interface {
	New(connectionString string, database string)
	GetDB() *MongoDbDatabase
}
