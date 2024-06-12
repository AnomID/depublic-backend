package binder

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserCreateRequest struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Fullname     string `json:"fullname" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Role         string `json:"role" `
	Status       bool   `json:"status" `
	Verification bool   `json:"verification" `
}

type UserUpdateRequest struct {
	User_ID      string `param:"user_id" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Fullname     string `json:"fullname" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Role         string `json:"role" `
	Status       bool   `json:"status" `
	Verification bool   `json:"verification" `
}

type UserDeleteRequest struct {
	User_ID string `param:"user_id" validate:"required"`
}
