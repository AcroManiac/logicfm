package main

import (
	"context"
	"log"

	"github.com/ahamtat/logicfm/internal/adapters/logger"
	"github.com/ahamtat/logicfm/internal/adapters/swagger"
	"github.com/ahamtat/logicfm/internal/adapters/swagger/builder/restapi"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(logger.New),
		fx.Provide(swagger.NewBuilder),
		fx.Invoke(register),
	).Run()
}

func register(lc fx.Lifecycle, server *restapi.Server) {
	// Append start and shutdown hooks for server object
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.Serve(); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown()
		},
	})
}
