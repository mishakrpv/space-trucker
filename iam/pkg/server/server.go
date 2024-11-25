package server

import (
	"context"
	"net/http"
	"sync"

	"github.com/jmoiron/sqlx"
)

func New() *Server {
	return &Server{}
}

type Server struct {
	ctx              context.Context
	shutdownFn       context.CancelFunc
	shutdownOnce     sync.Once
	shutdownFinished chan struct{}
	mtx              sync.Mutex

	db         sqlx.DB
	HTTPServer *http.Server
}
