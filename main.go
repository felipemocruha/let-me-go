package main

import (
	"github.com/rs/zerolog/log"

	"github.com/felipemocruha/let-me-go/config"
	"github.com/felipemocruha/let-me-go/db"
	"go.uber.org/fx"
)

func main() {
	log.Logger = log.With().Caller().Logger()
	
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			db.NewDatabase,
		),
		//fx.Invoke(Register),
	)

	app.Run()
}
