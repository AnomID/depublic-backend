package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	User_ID            uuid.UUID `json:"user_id"`
	Fullname           string    `json:"fullname"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Phone              string    `json:"phone"`
	Role               string    `json:"role"`
	Status             bool      `json:"status"`
	ResetCode          string    `json:"reset_code"`
	ResetCodeExpiresAt time.Time `json:"reset_code_expires_at"`
	Auditable
	Verification     bool   `json:"verification"`
	VerificationCode string `json:"verification_code"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Role == "" {
		u.Role = "user"
	}
	if !u.Status {
		u.Status = true
	}
	return
}

func NewUser(fullname, email, password, phone, role string, status, verification bool) *User {
	return &User{
		User_ID:      uuid.New(),
		Fullname:     fullname,
		Email:        email,
		Password:     password,
		Phone:        phone,
		Role:         role,
		Status:       status,
		Verification: verification,
		Auditable:    NewAuditable(),
	}
}

func UpdateUser(user_id uuid.UUID, fullname, email, password, phone, role string, status, verification bool) *User {
	return &User{
		User_ID:      user_id,
		Fullname:     fullname,
		Email:        email,
		Password:     password,
		Phone:        phone,
		Role:         role,
		Status:       status,
		Verification: verification,
		Auditable:    UpdateAuditable(),
	}
}
