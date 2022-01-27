package authdto

type Register struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Username string `form:"username" json:"username" binding:"required" `
	Password string `form:"password" json:"password" binding:"required" `
}
