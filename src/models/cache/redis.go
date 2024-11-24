package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-gorm/caches/v4"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	rdb *redis.Client
}

func (r *RedisCache) GetCachePlugin() *caches.Caches {
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return &caches.Caches{
		Conf: &caches.Config{
			Cacher: r,
		},
	}
}

func (r *RedisCache) Get(ctx context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	res, err := r.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if err := q.Unmarshal([]byte(res)); err != nil {
		return nil, err
	}

	return q, nil
}

func (r *RedisCache) Store(ctx context.Context, key string, val *caches.Query[any]) error {
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	r.rdb.Set(ctx, key, res, 300*time.Second)
	return nil
}

func (r *RedisCache) Invalidate(ctx context.Context) error {
	var (
		cursor uint64
		keys   []string
	)
	for {
		var (
			k   []string
			err error
		)
		k, cursor, err = r.rdb.Scan(ctx, cursor, fmt.Sprintf("%s*", caches.IdentifierPrefix), 0).Result()
		if err != nil {
			return err
		}
		keys = append(keys, k...)
		if cursor == 0 {
			break
		}
	}

	if len(keys) > 0 {
		if _, err := r.rdb.Del(ctx, keys...).Result(); err != nil {
			return err
		}
	}
	return nil
}
