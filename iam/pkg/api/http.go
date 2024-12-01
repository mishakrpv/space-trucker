package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	services := &Services{
		cfg:                    cfg,
		pgConnectionFactory:    postgres.NewConnectionFactory(cfg.Postgres),
		redisConnectionFactory: redis.NewConnectionFactory(cfg.Redis),
	}

	return services
}

type httpHandler struct {
	engine *gin.Engine

	services *Services
}

func NewHTTPHandler(services *Services) http.Handler {
	engine := gin.New()
	engine.Use(gin.Recovery())

	h := &httpHandler{
		engine:   engine,
		services: services,
	}

	h.setupRouters()

	return h.engine
}

func (h *httpHandler) setupRouters() {
	h.engine.POST("/login", h.login)
	h.engine.POST("/signup", h.signup)
}

func (h *httpHandler) login(c *gin.Context) {

}

func (h *httpHandler) signup(c *gin.Context) {

}
