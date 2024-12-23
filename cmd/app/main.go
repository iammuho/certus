package main

import (
	"github.com/iammuho/certus/cmd/app/config"
	"github.com/iammuho/certus/cmd/app/context"
	"github.com/iammuho/certus/internal/hub"
	"github.com/iammuho/certus/pkg/logger"

	_ "github.com/iammuho/certus/internal/drivers/aws"
)

func main() {
	l, err := logger.NewLogger(
		logger.WithLoggerLevel(config.Config.Logger.Level),
		logger.WithLoggerName(config.Config.Logger.Name),
	)

	if err != nil {
		panic(err)
	}

	// Create context
	ctx := context.NewAppContext(l)

	// Initialize the hub system
	h := hub.NewHub(ctx)

	// Execute all drivers
	h.ExecuteDrivers()
}
