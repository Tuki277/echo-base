package model

import (
	"echo-base/internal/contract"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FullName     string
	Email        string
	Password     string
	RoleId       uint
	StatusUserId uint
}

func (m *User) ToContract() *contract.UserResponse {
	if m == nil {
		return &contract.UserResponse{}
	}

	result := &contract.UserResponse{
		FullName:  m.FullName,
		Email:     m.Email,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
	}
	return result
}
