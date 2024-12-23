package context

import (
	contextPKG "context"

	"github.com/iammuho/certus/pkg/logger"
)

//go:generate mockgen -destination=mocks/mock_app_contexter.go -package=mockcontext -source=context.go

type AppContext interface {
	GetContext() contextPKG.Context
	GetContextWithTimeout() (contextPKG.Context, contextPKG.CancelFunc)

	// Custom Packages
	GetLogger() *logger.Logger
}
