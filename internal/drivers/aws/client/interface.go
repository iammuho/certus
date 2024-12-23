package aws

import "github.com/aws/aws-sdk-go-v2/service/s3"

type AWSContext interface {
	GetS3Client() *s3.Client
}
