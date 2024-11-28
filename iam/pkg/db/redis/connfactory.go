package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/space-trucker/iam/pkg/config"
)

type ConnectionFactory interface {
	Create(ctx context.Context) *redis.Client
}

type connectionFactory struct {
	redisCfg *config.Redis
}

func NewConnectionFactory() ConnectionFactory {
	return &connectionFactory{}
}

func (f *connectionFactory) Create(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(f.redisCfg.CreateRedisOptions())

	go func(ctx context.Context) {
		<-ctx.Done()
		logger := log.Ctx(ctx)
		logger.Info().Msg("Closing Redis connection...")
		err := rdb.Close()
		if err != nil {
			logger.Error().Err(err).
				Msg("An error occured while closing Redis connection")
		}
	}(ctx)

	return rdb
}
