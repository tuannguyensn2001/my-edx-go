package authdto

type Register struct {
	Email    string `form:"email" json:"email" binding:"required" validate:"required,email"`
	Username string `form:"username" json:"username" binding:"required" validate:"required" `
	Password string `form:"password" json:"password" binding:"required" validate:"required" `
}
