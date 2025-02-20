package consts

type UserEntity struct {
	Id           int     `json:"id" db:"id"`
	Name         string  `json:"name" db:"name" validate:"required"`
	Email        string  `json:"email" db:"email" validate:"required,email"`
	Password     string  `json:"password" db:"password" validate:"required,min=8"`
	Status       string  `json:"status" db:"status"`
	CreatedAt    string  `json:"created_at" db:"created_at"`
	UpdatedAt    string  `json:"updated_at" db:"updated_at"`
	ReFreshToken *string `json:"refresh_token" db:"refresh_token"`
}

type RegisterRequest struct {
	UserEntity
	ConfirmPassword string `json:"confirm_password"  validate:"required,eqfield=Password"`
}

type UserResponse struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Status    string `json:"status" db:"status"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,min=8"`
}

type LoginResponse struct {
	UserResponse
	Token        string `json:"token"`
	ReFreshToken string `json:"refresh_token"`
}
