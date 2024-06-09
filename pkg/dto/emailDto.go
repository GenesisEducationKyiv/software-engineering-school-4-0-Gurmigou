package dto

type EmailDto struct {
	Email string `form:"email" binding:"required,email"`
}
