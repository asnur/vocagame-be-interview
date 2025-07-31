package shared

type (
	CheckPasswordRequest struct {
		HashedPassword string `json:"input"`
		Password       string `json:"key"`
	}

	CheckPasswordResponse struct {
		IsValid bool `json:"is_valid"`
	}
)
