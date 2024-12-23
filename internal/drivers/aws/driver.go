package aws

import (
	"os"

	"github.com/cucumber/godog"
	"github.com/iammuho/certus/cmd/app/context"
	"github.com/iammuho/certus/internal/drivers/aws/config"
	"github.com/iammuho/certus/internal/drivers/aws/providers"
	"github.com/iammuho/certus/internal/hub"
)

// AWSDriver is the driver for AWS
type AWSDriver struct{}

func init() {
	hub.Register("aws", &AWSDriver{})
}

func (l *AWSDriver) Execute(ctx context.AppContext) {
	ctx.GetLogger().Info("Executing AWS driver")

	// // Initialize AWS Client
	// aws, err := aws.NewClient(
	// 	aws.WithAWSRegion(config.Config.AWS.Region),
	// 	aws.WithAWSAccessKeyID(config.Config.AWS.AccessKeyID),
	// 	aws.WithAWSSecretAccessKey(config.Config.AWS.SecretAccessKey),
	// )

	// if err != nil {
	// 	ctx.GetLogger().Error("Error initializing AWS client", zap.Error(err))
	// 	return
	// }

	// Define options for godog
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{config.Config.Application.FeaturesPath},
	}
	// Create and run the test suite
	status := godog.TestSuite{
		Name: "aws",
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			// Initialize the providers
			providers.InitializeSecurityGroupScenario(ctx)
		},
		Options: &opts,
	}.Run()

	// Exit with the status code
	os.Exit(status)
}
