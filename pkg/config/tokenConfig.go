package config

import (
	"time"

	"github.com/asnur/vocagame-be-interview/pkg/utils"
)

type TokenConfig struct {
	KeyId           string        `envconfig:"JWK_KID" required:"true"`
	SignKey         string        `envconfig:"ACCESS_TOKEN_RSA256_PRIVATE_KEY" required:"true"` //RSA Private Key in PEM
	VerifyKey       string        `envconfig:"ACCESS_TOKEN_RSA256_PUBLIC_KEY" required:"true"`  //RSA Public Key in PEM
	IatLeeway       time.Duration `envconfig:"TOKEN_IAT_LEEWAY" required:"true"`                //Leeway time for iat to accommodate server time discrepancy
	Audiences       []string      `envconfig:"TOKEN_AUDIENCES" required:"true"`
	Issuer          string        `envconfig:"TOKEN_ISSUER" required:"true"`
	AccessValidFor  time.Duration `envconfig:"ACCESS_TOKEN_VALID_FOR" required:"true"`
	RefreshValidFor time.Duration `envconfig:"REFRESH_TOKEN_VALID_FOR" required:"true"`
}

func NewTokenConfig() (TokenConfig, error) {
	var config TokenConfig

	if err := utils.NewConfig(&config); err != nil {
		return TokenConfig{}, err
	}

	return config, nil
}
