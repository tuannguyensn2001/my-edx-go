package dto

type LoginDTO struct {
	Email string `form:"email" json:"email"  binding:"required"`
}
