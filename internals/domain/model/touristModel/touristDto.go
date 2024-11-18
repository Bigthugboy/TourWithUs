package touristModel

type RegisterRequest struct {
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
	ProfilePic string `json:"profilePic"`
	Username   string `json:"username"validate:"required"`
}

type RegisterResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName" `
	Email     string `json:"email" `
	Username  string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
type LoginResponse struct {
	Username    string `json:"username"`
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
}
