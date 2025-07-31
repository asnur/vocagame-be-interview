package injection

import (
	"crypto/rsa"

	"github.com/asnur/vocagame-be-interview/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	RsaPrivateKey *rsa.PrivateKey
	RsaPublicKey  *rsa.PublicKey
}

func NewJwt(config config.TokenConfig) (Jwt, error) {
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.SignKey))
	if err != nil {
		return Jwt{}, err
	}
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.VerifyKey))
	if err != nil {
		return Jwt{}, err
	}
	return Jwt{
		RsaPrivateKey: rsaPrivateKey,
		RsaPublicKey:  rsaPublicKey,
	}, nil
}
