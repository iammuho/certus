package main

import (
	"context"

	"github.com/iammuho/certus/cmd/app/config"
	"github.com/iammuho/certus/pkg/aws"
	"github.com/iammuho/certus/pkg/logger"
)

func main() {
	l, err := logger.NewLogger(
		logger.WithLoggerLevel(config.Config.Logger.Level),
		logger.WithLoggerName(config.Config.Logger.Name),
	)

	if err != nil {
		panic(err)
	}

	// AWS
	aws, err := aws.NewClient(
		aws.WithAWSRegion(config.Config.AWS.Region),
		aws.WithAWSAccessKeyID(config.Config.AWS.AccessKeyID),
		aws.WithAWSSecretAccessKey(config.Config.AWS.SecretAccessKey),
	)

	if err != nil {
		panic(err)
	}

	// S3
	s3Client := aws.GetS3Client()

	// S3 Bucket
	buckets, err := s3Client.ListBuckets(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	for _, bucket := range buckets.Buckets {
		l.Info(*bucket.Name)
	}
}
