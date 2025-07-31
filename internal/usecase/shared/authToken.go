package shared

import (
	"context"
	"fmt"
	"time"

	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/shared"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

const (
	TokenType = "Bearer"
)

func (t *usecase) AuthToken(ctx context.Context, req ucModel.AuthTokenRequest) (ucModel.AuthTokenResponse, error) {
	var (
		response ucModel.AuthTokenResponse
	)

	accessTokenStr, accessToken, err := t.generateAccessToken(req)
	if err != nil {
		t.resource.Logger.WithContext(ctx).Errorf("Error generate access token: %v", err)
		return response, err
	}

	refreshTokenStr, err := t.generateRefreshToken(*accessToken)
	if err != nil {
		t.resource.Logger.WithContext(ctx).Errorf("Error generate refresh token: %v", err)
		return response, err
	}

	response = ucModel.AuthTokenResponse{
		Token: oauth2.Token{
			AccessToken:  accessTokenStr,
			RefreshToken: refreshTokenStr,
			TokenType:    TokenType,
			Expiry:       accessToken.ExpiresAt.Time,
		},
	}

	return response, nil
}

func (t *usecase) generateAccessToken(req ucModel.AuthTokenRequest) (string, *ucModel.AccessToken, error) {
	timeNow := time.Now()
	accessToken := ucModel.AccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(timeNow.Add(t.resource.TokenConfig.AccessValidFor)),
			IssuedAt:  jwt.NewNumericDate(timeNow.Add(t.resource.TokenConfig.IatLeeway * -1)),
			Issuer:    t.resource.TokenConfig.Issuer,
			Audience:  jwt.ClaimStrings(t.resource.TokenConfig.Audiences),
			Subject:   fmt.Sprintf("%d", req.UserId),
		},
		Data: ucModel.AccessTokenClaimData{
			UserId: req.UserId,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessToken)
	jwtToken.Header["kid"] = t.resource.TokenConfig.KeyId
	accessTokenString, err := jwtToken.SignedString(t.resource.Jwt.RsaPrivateKey)
	if err != nil {
		return "", nil, err
	}

	return accessTokenString, &accessToken, nil
}

func (s *usecase) generateRefreshToken(accessToken ucModel.AccessToken) (string, error) {
	timeNow := time.Now()
	refreshToken := ucModel.RefreshToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(timeNow.Add(s.resource.TokenConfig.RefreshValidFor)),
			IssuedAt:  jwt.NewNumericDate(timeNow.Add(s.resource.TokenConfig.IatLeeway * -1)),
			Issuer:    accessToken.Issuer,
			Audience:  jwt.ClaimStrings(accessToken.Audience),
			Subject:   accessToken.Subject,
		},
		Role: accessToken.Role,
		Data: ucModel.RefreshTokenClaimData{
			AccessTokenId: accessToken.ID,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshToken)
	refreshTokenString, err := jwtToken.SignedString(s.resource.Jwt.RsaPrivateKey)
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

type IAuthToken interface {
	AuthToken(ctx context.Context, req ucModel.AuthTokenRequest) (ucModel.AuthTokenResponse, error)
}
