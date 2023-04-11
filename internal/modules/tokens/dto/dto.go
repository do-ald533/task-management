package dto

type TokenResponse struct {
	Token       string   `json:"token"`
	Credentials []string `json:"credentials"`
	Role        string   `json:"role"`
}

type TokenRequest struct {
	Secret string `json:"secret"`
	Role   string `json:"role"`
}
