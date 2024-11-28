package config

import "github.com/redis/go-redis/v9"

type Redis struct {
}

func (r *Redis) CreateRedisOptions() *redis.Options {
	return &redis.Options{}
}
