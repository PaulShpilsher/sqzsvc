package auth

type CredentialsInput struct {
	Email    string `json:"email" binding:"required" validate:"required,email,max=256"`
	Password string `json:"password" binding:"required" validate:"required,min=6,max=32"`
}
