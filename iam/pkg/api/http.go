package api

import (
	"net/http"

	"github.com/space-trucker/iam/pkg/config"
	"github.com/space-trucker/iam/pkg/db/postgres"
	"github.com/space-trucker/iam/pkg/db/redis"
)

type Services struct {
	cfg *config.Cfg

	pgConnectionFactory    postgres.ConnectionFactory
	redisConnectionFactory redis.ConnectionFactory
}

func NewServices(cfg *config.Cfg) *Services {
	services := &Services{}

	return services
}

func NewHTTPHandler(services *Services) http.Handler {
	return nil
}
