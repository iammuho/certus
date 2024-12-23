package aws

import (
	"github.com/iammuho/certus/cmd/app/context"
	aws "github.com/iammuho/certus/internal/drivers/aws/client"
	"github.com/iammuho/certus/internal/drivers/aws/config"
	"github.com/iammuho/certus/internal/hub"
	"go.uber.org/zap"
)

// AWSDriver is the driver for AWS
type AWSDriver struct{}

func init() {
	hub.Register("aws", &AWSDriver{})
}

func (l *AWSDriver) Execute(ctx context.AppContext) {
	ctx.GetLogger().Info("Executing AWS driver")

	// Initialize AWS Client
	aws, err := aws.NewClient(
		aws.WithAWSRegion(config.Config.AWS.Region),
		aws.WithAWSAccessKeyID(config.Config.AWS.AccessKeyID),
		aws.WithAWSSecretAccessKey(config.Config.AWS.SecretAccessKey),
	)

	if err != nil {
		ctx.GetLogger().Error("Error initializing AWS client", zap.Error(err))
		return
	}

	// S3
	s3Client := aws.GetS3Client()

	// S3 Bucket
	buckets, err := s3Client.ListBuckets(ctx.GetContext(), nil)

	if err != nil {
		panic(err)
	}

	for _, bucket := range buckets.Buckets {
		ctx.GetLogger().Info("Bucket", zap.String("name", *bucket.Name))
	}
}
