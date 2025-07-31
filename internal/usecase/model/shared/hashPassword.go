package shared

type (
	HashPasswordRequest struct {
		Input string `json:"input"`
		Key   string `json:"key"`
	}

	HashPasswordResponse struct {
		Hash string `json:"hash"`
	}
)
