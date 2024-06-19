package subscriber

type EmailDto struct {
	Email string `form:"email" binding:"required,email"`
}
