package main

import (
	"context"
	
	"github.com/felipemocruha/let-me-go/config"
	"github.com/felipemocruha/let-me-go/logging"	
	"github.com/felipemocruha/let-me-go/db"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		fx.Provide(config.LoadConfig),
		fx.Provide(logging.NewLogger),
		fx.Provide(db.NewDatabase),
		fx.Invoke(Start),
		fx.Logger(&log.Logger),
	)

	app.Run()
}

func Start(lc fx.Lifecycle, logger *zap.SugaredLogger, database db.Database) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("passou pela main")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			defer logger.Sync()
			return database.Close()
		},
	})	

}
