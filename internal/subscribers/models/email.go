package models

type EmailDto struct {
	Email string `form:"email" binding:"required,email"`
}
