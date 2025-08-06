package response

type SignInResponse struct {
	AccessToken string  `json:"access_token"`
	Role        string  `json:"role"`
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}
