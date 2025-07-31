package shared

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

type (
	AuthTokenRequest struct {
		UserId int64 `json:"user_id"`
	}

	AuthTokenResponse struct {
		oauth2.Token
	}

	ValidateRefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	ValidateRefreshTokenResponse struct {
		oauth2.Token
	}

	AccessToken struct {
		jwt.RegisteredClaims
		Role string               `json:"roles,omitempty"`
		Data AccessTokenClaimData `json:"dat"`
	}

	AccessTokenClaimData struct {
		UserId int64 `json:"user_id"`
		//additional future information here
	}

	RefreshToken struct {
		jwt.RegisteredClaims
		Role string                `json:"roles,omitempty"`
		Data RefreshTokenClaimData `json:"dat"`
	}

	RefreshTokenClaimData struct {
		AccessTokenId string `json:"token_id,omitempty"`
		//additional future information here
	}
)
