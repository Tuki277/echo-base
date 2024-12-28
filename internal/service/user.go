package service

import (
	c "echo-base/internal/contract"
	"echo-base/internal/dto"
	"echo-base/internal/model"
	"echo-base/internal/repository"
)

type UserService struct {
	repo repository.ModelAuthRepository[model.User, int, *c.UserRequest]
}

func NewUserService(UserRepo repository.ModelAuthRepository[model.User, int, *c.UserRequest]) *UserService {
	return &UserService{repo: UserRepo}
}

func (s *UserService) List() ([]*c.UserResponse, error) {
	results, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	var list []*c.UserResponse

	for _, result := range results {
		list = append(list, dto.UserModelToContract(result))
	}

	return list, nil
}

func (s *UserService) ChangeStatus(r *c.ChangeStatusRes) error {
	_, err := s.repo.GetById(r.Id)
	if err != nil {
		return err
	}
	err = s.repo.ChangeStatus(r)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Delete(ids []uint) error {

	err := s.repo.DeleteAList(ids)
	if err != nil {
		return err
	}

	return nil
}
