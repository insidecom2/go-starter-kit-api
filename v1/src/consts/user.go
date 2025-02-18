package consts

type UserUpdateRequest struct {
	Name     string `json:"name" db:"name" validate:"required"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min=8"`
	Status   string `json:"status" db:"status"`
}
