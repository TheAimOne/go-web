package auth

const (
	AUTH_TYPE_MOBILE = "mobile"
	AUTH_TYPE_EMAIL  = "email"
)

type AuthRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type AuthResponse struct {
	Data *AuthData `json:"data"`
}

type AuthData struct {
	Token string `json:"token"`
}
