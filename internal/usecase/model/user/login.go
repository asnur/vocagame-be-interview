package user

import obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"

type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

func (u LoginRequest) ToObModel() obModel.Users {
	return obModel.Users{
		Username: u.Username,
	}
}
