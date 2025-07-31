package config

import "github.com/asnur/vocagame-be-interview/pkg/utils"

type (
	AppConfig struct {
		AppName  string `envconfig:"APP_NAME" required:"true"`
		Version  string `envconfig:"APP_VERSION" required:"true"`
		Key      string `envconfig:"APP_KEY" required:"true"`
		Currency string `envconfig:"APP_CURRENCY_DEFAULT" default:"USD"`
	}
)

func NewAppConfig() (AppConfig, error) {
	var config AppConfig

	if err := utils.NewConfig(&config); err != nil {
		return AppConfig{}, err
	}

	return config, nil
}
