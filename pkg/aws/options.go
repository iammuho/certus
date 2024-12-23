package aws

// Option is the func interface to assign options
type Option func(*AWSOptions)

type AWSOptions struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
}

// WithAWSRegion defines the aws region to be used on application
func WithAWSRegion(region string) Option {
	return func(o *AWSOptions) {
		o.Region = region
	}
}

// WithAWSAccessKeyID defines the aws access key id to be used on application
func WithAWSAccessKeyID(accessKeyID string) Option {
	return func(o *AWSOptions) {
		o.AccessKeyID = accessKeyID
	}
}

// WithAWSSecretAccessKey defines the aws secret access key to be used on application
func WithAWSSecretAccessKey(secretAccessKey string) Option {
	return func(o *AWSOptions) {
		o.SecretAccessKey = secretAccessKey
	}
}
