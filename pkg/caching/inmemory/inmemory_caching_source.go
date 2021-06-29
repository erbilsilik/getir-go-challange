package inmemory

import (
	"github.com/dgraph-io/ristretto"
	"github.com/erbilsilik/getir-go-challange/pkg/caching/inmemory/interfaces"
)

type inmemoryCachingSource struct {
	cache *ristretto.Cache
}

func (i inmemoryCachingSource) DeleteValueByKey(key string) {
	i.cache.Del(key)
}

func (i inmemoryCachingSource) GetValueByKey(key string) interface{} {
	data, exists := i.cache.Get(key)
	if exists {
		return data
	}
	return nil
}

func (i inmemoryCachingSource) SetValue(key string, value interface{}) error {
	i.cache.Set(key, value, 1)
	return nil
}

func New() (interfaces.InMemoryCachingSource, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	if err != nil {
		return nil, err
	}

	return &inmemoryCachingSource{
		cache: cache,
	}, nil
}