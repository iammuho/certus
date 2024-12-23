// Package config aims to provide a config package
package config

import (
	"github.com/caarlos0/env"
)

func init() {
	Config = &config{}

	if err := env.Parse(&Config.AWS); err != nil {
		panic(err)
	}
}
