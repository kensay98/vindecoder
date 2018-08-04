package storage

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisStorage struct {
	client *redis.Client
}

func (r *RedisStorage) Get(key string) (value string) {
	value, _ = r.client.Get(key).Result()
	return
}

func (r *RedisStorage) Set(key, value string) {
	log.Info("writing to log")
	err := r.client.Set(key, value, time.Second * 1000).Err()
	if err != nil {
		log.Error(err)
	}
	return
}


func NewRedisStorage(host, password string, db int) *RedisStorage {
	return &RedisStorage{
		client: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       db,
		}),
	}
}
