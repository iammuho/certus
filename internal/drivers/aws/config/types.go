package config

// Config stores configuration values
var Config *config

type config struct {
	Application struct {
		FeaturesPath string `env:"APPLICATION_FEATURES_PATH" envDefault:"/Users/muhammetarslan/Projects/iammuho/certus/features/aws"`
	}

	AWS struct {
		Region          string `env:"AWS_REGION" envDefault:"eu-central-1"`
		AccessKeyID     string `env:"AWS_ACCESS_KEY_ID" envDefault:""`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY" envDefault:""`
	}
}
