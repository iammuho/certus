package context

import (
	"context"

	"github.com/iammuho/certus/pkg/logger"
)

type appContext struct {
	ctx    context.Context
	logger *logger.Logger
}

func NewAppContext(logger *logger.Logger) AppContext {
	ctx := context.Background()

	return &appContext{
		ctx:    ctx,
		logger: logger,
	}
}

func (c *appContext) GetContext() context.Context {
	return c.ctx
}

func (c *appContext) GetContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.ctx, 5)
}

func (c *appContext) GetLogger() *logger.Logger {
	return c.logger
}
