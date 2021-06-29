package interfaces

type InMemoryCachingSource interface {
	GetValueByKey(key string) interface{}
	SetValue(key string, value interface{}) error
	DeleteValueByKey(key string)
}
