package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Token    string `json:"token"`
}
