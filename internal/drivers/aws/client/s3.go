package aws

import "github.com/aws/aws-sdk-go-v2/service/s3"

// GetS3Client returns the S3 client
func (c *client) GetS3Client() *s3.Client {
	return c.S3Client
}
