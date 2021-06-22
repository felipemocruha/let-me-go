package db

import (
	"fmt"
	"context"
	
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/felipemocruha/let-me-go/config"
	"go.uber.org/fx"
)

type Postgres struct {
	conn *sqlx.DB
}

func (db Postgres) Close() error {
	return db.conn.Close()
}

type Database interface {
	Close() error
}

func NewDatabase(lc fx.Lifecycle, config config.Config) Database {
	conn, err := sqlx.Connect("postgres", makeConnStr(config))
	if err != nil {
		log.Fatal().Msgf("failed to create database connection: %v", err)
	}
	postgres := &Postgres{conn}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return postgres.Close()
		},
	})	

	return postgres
}

func makeConnStr(config config.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		config.Database.Host,
		config.Database.User,
		config.Database.Name,
		config.Database.Password,
		config.Database.Port,
	)
}
