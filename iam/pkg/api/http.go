package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"github.com/space-trucker/iam/pkg/config"
	"github.com/space-trucker/iam/pkg/db/postgres"
	"github.com/space-trucker/iam/pkg/db/redis"
)

var jwtKey = []byte("my_secret_key")

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

type Credentials struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (h *httpHandler) login(c *gin.Context) {
	var creds Credentials
	if err := c.Bind(&creds); err != nil {
		log.Info().Err(err).Msg("Failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(&creds)
	if err != nil {
		log.Info().Err(err).Msg("Request validation failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// conn, err := h.services.pgConnectionFactory.Create(c)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Failed to open DB connection")
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured processing the request"})
	// }

	// query := `SELECT * FROM user;`

	// _, err = conn.DB.Query(query)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Failed to execute query")
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured processing the request"})
	// }

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &claims{
		Email: creds.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Error().Err(err).Msg("Failed to sign token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured processing the request"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *httpHandler) signup(c *gin.Context) {

}
