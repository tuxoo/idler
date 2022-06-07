package dto

type UserDTO struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type SignInDTO struct {
	Email    string `json:"email" binding:"required,email,max=64" example:"kill-77@mail.ru"`
	Password string `json:"password" binding:"required,min=6,max=64" example:"qwerty"`
}

type SignUpDTO struct {
	Name     string `json:"name" binding:"required,min=2,max=64" example:"alex"`
	Email    string `json:"email" binding:"required,email,max=64" example:"kill-77@mail.ru"`
	Password string `json:"password" binding:"required,min=6,max=64" example:"qwerty"`
}
