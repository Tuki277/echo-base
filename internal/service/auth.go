package service

import (
	c "echo-base/internal/contract"
	"echo-base/internal/model"
	"echo-base/internal/repository"
	"echo-base/utils/constants"
	"echo-base/utils/hash_password"
	"errors"
	"gorm.io/gorm"
)

type AuthService struct {
	repo repository.ModelAuthRepository[model.User, int, *c.UserRequest]
}

const (
	Salt   = "REC"
	JWTkey = "REC"
)

func NewAuthService(jobRepo repository.ModelAuthRepository[model.User, int, *c.UserRequest]) *AuthService {
	return &AuthService{repo: jobRepo}
}

func (s *AuthService) Register(r *c.RegisterRequest) (*c.AuthResponse, error) {
	var userModel model.User
	user, err := s.repo.GetByEmail(r.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if user != nil {
		return nil, errors.New(constants.ErrEmailExisting)
	}
	hashPassword, err := hash_password.HashPassword(r.Password + Salt)
	if err != nil {
		return nil, err
	}
	userModel.FullName = r.FullName
	userModel.Email = r.Email
	userModel.Password = hashPassword
	err = s.repo.Add(&userModel)
	if err != nil {
		return nil, err
	}
	return &c.AuthResponse{
		IsSuccessfull: true,
	}, nil
}

func (s *AuthService) Login(r *c.LoginRequest) (*c.AuthResponse, error) {
	user, err := s.repo.GetByEmail(r.Email)
	if err != nil {
		return nil, err
	}
	isCheck := hash_password.CheckPasswordHash(r.Password+Salt, user.Password)
	if !isCheck {
		return nil, errors.New("info not match")
	}
	j := hash_password.NewJWT(JWTkey)
	token, err := j.GenerateToken(user.ID, user.FullName, user.Email, user.StatusUserId, user.RoleId)
	if err != nil {
		return nil, err
	}
	return &c.AuthResponse{
		IsSuccessfull: true,
		Token:         token,
	}, nil
}
