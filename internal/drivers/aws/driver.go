package aws

import (
	"github.com/iammuho/certus/cmd/app/context"
	"github.com/iammuho/certus/internal/hub"
)

// AWSDriver is a dummy implementation of the Driver interface.
type AWSDriver struct {
	isActive bool
}

func init() {
	hub.Register("aws", &AWSDriver{})
}

func (l *AWSDriver) Activate() error {
	l.isActive = true
	return nil
}

func (l *AWSDriver) Deactivate() error {
	l.isActive = false
	return nil
}

func (l *AWSDriver) Execute(ctx context.AppContext) {
	ctx.GetLogger().Info("I'm the AWS driver")
}
