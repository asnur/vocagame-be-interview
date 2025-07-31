package shared

import (
	"context"
	"errors"
	"fmt"

	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/shared"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

func (u *usecase) ValidateToken(ctx context.Context, token string) (oauth2.Token, error) {
	// Parse and validate the JWT token
	jwtToken, err := jwt.ParseWithClaims(token, &ucModel.AccessToken{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Check if the key ID matches
		if kid, ok := token.Header["kid"].(string); ok {
			if kid != u.resource.TokenConfig.KeyId {
				return nil, fmt.Errorf("invalid key ID: %v", kid)
			}
		} else {
			return nil, fmt.Errorf("missing key ID in token header")
		}

		return u.resource.Jwt.RsaPublicKey, nil
	})

	if err != nil {
		u.resource.Logger.WithContext(ctx).Errorf("Error parsing token: %v", err)
		return oauth2.Token{}, fmt.Errorf("invalid token: %w", err)
	}

	// Check if token is valid
	if !jwtToken.Valid {
		u.resource.Logger.WithContext(ctx).Error("Token is not valid")
		return oauth2.Token{}, errors.New("token is not valid")
	}

	// Extract claims
	claims, ok := jwtToken.Claims.(*ucModel.AccessToken)
	if !ok {
		u.resource.Logger.WithContext(ctx).Error("Invalid token claims")
		return oauth2.Token{}, errors.New("invalid token claims")
	}

	// Validate issuer
	if claims.Issuer != u.resource.TokenConfig.Issuer {
		u.resource.Logger.WithContext(ctx).Errorf("Invalid issuer: %v", claims.Issuer)
		return oauth2.Token{}, fmt.Errorf("invalid issuer: %v", claims.Issuer)
	}

	// Validate audience
	if len(claims.Audience) == 0 {
		u.resource.Logger.WithContext(ctx).Error("Token has no audience")
		return oauth2.Token{}, errors.New("token has no audience")
	}

	validAudience := false
	for _, tokenAud := range claims.Audience {
		for _, configAud := range u.resource.TokenConfig.Audiences {
			if tokenAud == configAud {
				validAudience = true
				break
			}
		}
		if validAudience {
			break
		}
	}

	if !validAudience {
		u.resource.Logger.WithContext(ctx).Errorf("Invalid audience: %v", claims.Audience)
		return oauth2.Token{}, fmt.Errorf("invalid audience: %v", claims.Audience)
	}

	// Create oauth2.Token from the validated claims
	oauthToken := oauth2.Token{
		AccessToken: token,
		TokenType:   TokenType,
		Expiry:      claims.ExpiresAt.Time,
	}

	u.resource.Logger.WithContext(ctx).Infof("Token validated successfully for user: %d", claims.Data.UserId)

	return oauthToken, nil
}

// ValidateTokenAndGetClaims validates token and returns the claims
func (u *usecase) ValidateTokenAndGetClaims(ctx context.Context, token string) (*ucModel.AccessToken, error) {
	// Parse and validate the JWT token
	jwtToken, err := jwt.ParseWithClaims(token, &ucModel.AccessToken{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Check if the key ID matches
		if kid, ok := token.Header["kid"].(string); ok {
			if kid != u.resource.TokenConfig.KeyId {
				return nil, fmt.Errorf("invalid key ID: %v", kid)
			}
		} else {
			return nil, fmt.Errorf("missing key ID in token header")
		}

		return u.resource.Jwt.RsaPublicKey, nil
	})

	if err != nil {
		u.resource.Logger.WithContext(ctx).Errorf("Error parsing token: %v", err)
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Check if token is valid
	if !jwtToken.Valid {
		u.resource.Logger.WithContext(ctx).Error("Token is not valid")
		return nil, errors.New("token is not valid")
	}

	// Extract claims
	claims, ok := jwtToken.Claims.(*ucModel.AccessToken)
	if !ok {
		u.resource.Logger.WithContext(ctx).Error("Invalid token claims")
		return nil, errors.New("invalid token claims")
	}

	// Validate issuer
	if claims.Issuer != u.resource.TokenConfig.Issuer {
		u.resource.Logger.WithContext(ctx).Errorf("Invalid issuer: %v", claims.Issuer)
		return nil, fmt.Errorf("invalid issuer: %v", claims.Issuer)
	}

	// Validate audience
	if len(claims.Audience) == 0 {
		u.resource.Logger.WithContext(ctx).Error("Token has no audience")
		return nil, errors.New("token has no audience")
	}

	validAudience := false
	for _, tokenAud := range claims.Audience {
		for _, configAud := range u.resource.TokenConfig.Audiences {
			if tokenAud == configAud {
				validAudience = true
				break
			}
		}
		if validAudience {
			break
		}
	}

	if !validAudience {
		u.resource.Logger.WithContext(ctx).Errorf("Invalid audience: %v", claims.Audience)
		return nil, fmt.Errorf("invalid audience: %v", claims.Audience)
	}

	u.resource.Logger.WithContext(ctx).Infof("Token validated successfully for user: %d", claims.Data.UserId)

	return claims, nil
}

type IValidateToken interface {
	ValidateToken(ctx context.Context, token string) (oauth2.Token, error)
	ValidateTokenAndGetClaims(ctx context.Context, token string) (*ucModel.AccessToken, error)
}
