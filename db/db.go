package db

import (
	"fmt"
	
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/felipemocruha/let-me-go/config"
	"go.uber.org/zap"
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

func NewDatabase(config config.Config, logger *zap.SugaredLogger) (Database, error) {
	conn, err := sqlx.Connect("postgres", makeConnStr(config))
	if err != nil {
		logger.Error("database connection: %v", err)
		return nil, err
	}

	return &Postgres{conn}, nil
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
