package config

// Config stores configuration values
var Config *config

type config struct {

	// Application provides the application basic configurations.
	Application struct {
		Name        string `env:"APPLICATION_NAME"     envDefault:"Certus"`
		Environment string `env:"APPLICATION_ENV"     envDefault:"development"`
		Version     string `env:"APPLICATION_VERSION"     envDefault:"0.1"`
	}

	// Logger provides the logger basic configurations.
	Logger struct {
		Level string `env:"LOGGER_LEVEL"     envDefault:"debug"`
		Name  string `env:"LOGGER_NAME"     envDefault:"certus"`
	}

	// AWS
	AWS struct {
		Region          string `env:"AWS_REGION" envDefault:"eu-central-1"`
		AccessKeyID     string `env:"AWS_ACCESS_KEY_ID" envDefault:""`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY" envDefault:""`
	}
}
