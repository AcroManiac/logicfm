package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module is the logger module that can be passed into an Fx app.
var Module = fx.Provide(New)

// Params is the input parameter struct for the logger module.
type Params struct {
	fx.In

	LogLevel string
}

// New constructs a new logger.
func New() (*zap.SugaredLogger, error) {
	Init("debug")
	return sugar, nil
}
