package cache

import "errors"

type CacheFactory struct{}

func (f *CacheFactory) CreateCachePlugin(cacheType string) (Cache, error) {
	switch cacheType {
	case "redis":
		return &RedisCache{}, nil
	default:
		return nil, errors.New("cache type not supported")
	}
}
