package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/space-trucker/iam/pkg/config"
)

type ConnectionFactory interface {
	Create(ctx context.Context) (*sqlx.DB, error)
}

type connectionFactory struct {
	postgresCfg *config.Postgres
}

func NewConnectionFactory(postgresCfg *config.Postgres) ConnectionFactory {
	return &connectionFactory{
		postgresCfg: postgresCfg,
	}
}

func (f *connectionFactory) Create(ctx context.Context) (*sqlx.DB, error) {
	return open(ctx, f.postgresCfg)
}
