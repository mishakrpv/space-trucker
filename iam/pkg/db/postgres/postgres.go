package postgres

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

func open(ctx context.Context, postgresCfg *config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", postgresCfg.ConnectionString)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		logger := log.Ctx(ctx)
		logger.Info().Msg("Closing PostgreSQL connection...")
		err := db.Close()
		if err != nil {
			logger.Error().Err(err).
				Msg("An error occured while closing PostgreSQL connection")
		}
	}(ctx)

	db.MustExec(schema)

	return db, nil
}
