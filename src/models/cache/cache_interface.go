package cache

import "github.com/go-gorm/caches/v4"

type Cache interface {
	GetCachePlugin() *caches.Caches
}
