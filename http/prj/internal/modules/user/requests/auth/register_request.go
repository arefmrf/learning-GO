package auth

type RegisterRequest struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=32"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}
