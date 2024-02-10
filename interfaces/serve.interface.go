package interfaces

type ServeResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"serve_name"`
	Duration int16  `json:"serve_duration"`
	Color    string `json:"serve_color"`
	UserId   uint   `json:"user_id"`
}