package interfaces

type Response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}