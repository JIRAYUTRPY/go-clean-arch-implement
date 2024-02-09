package interfaces

type GetByIdRequest struct {
	ID uint `json:"user_id"`
}

type GetByEmailRequest struct {
	Email string `json:"email"`
}

type GetResponse struct {
	ID           uint
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	LocationLink string `json:"location_link"`
	StoreName    string `json:"store_name"`
	Role         string `json:"role"`
}