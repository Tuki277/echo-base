package dto

import (
	"echo-base/internal/contract"
	"echo-base/internal/model"
	"time"
)

func UserContractToModel(o *contract.UserRequest) *model.User {
	if o == nil {
		return &model.User{}
	}

	return &model.User{
		FullName: o.FullName,
		Email:    o.Email,
	}
}

func UserModelToContract(o *model.User) *contract.UserResponse {
	if o == nil {
		return &contract.UserResponse{}
	}

	result := &contract.UserResponse{
		FullName:  o.FullName,
		Email:     o.Email,
		CreatedAt: o.CreatedAt.Format(time.RFC3339),
		UpdatedAt: o.UpdatedAt.Format(time.RFC3339),
	}

	return result
}
