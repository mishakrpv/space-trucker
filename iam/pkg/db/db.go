package db

import (
	"context"
	_ "embed"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/space-trucker/iam/pkg/config"
)

var schema = `
CREATE TABLE IF NOT EXISTS user (
	user_id UUID PRIMARY KEY,
	email text NOT NULL,
	password_hash text NOT NULL
)
`

func Open(ctx context.Context, cfg config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.ConnectionString)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}
