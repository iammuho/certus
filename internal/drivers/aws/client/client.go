package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type client struct {
	Options  *AWSOptions
	S3Client *s3.Client
}

func NewClient(opts ...Option) (AWSContext, error) {
	options := AWSOptions{}
	for _, o := range opts {
		o(&options)
	}

	c := &client{
		Options: &options,
	}

	// Create S3Service
	c.createS3Client()

	return c, nil
}

// createS3Client attaches S3 client to the client
func (c *client) createS3Client() error {

	// Create S3Service
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading SDK config, ", err)
		return err
	}
	c.S3Client = s3.NewFromConfig(sdkConfig)

	return nil
}
