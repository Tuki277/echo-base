package repository

import (
	"echo-base/internal/contract"
	"echo-base/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ModelAuthRepository[model.User, int, *contract.UserRequest] {
	return &UserRepository{db: db}
}
func (r *UserRepository) DeleteAList(ids []uint) error {
	err := r.db.Table("users").Where("id IN ?", ids).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var userModel model.User
	err := r.db.Table("users").Where("email = ? AND deleted_at IS NULL", email).First(&userModel).Error
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}
func (r *UserRepository) GetById(id uint) (*model.User, error) {
	var userModel model.User
	err := r.db.Table("users").Where("id = ? AND deleted_at IS NULL", id).First(&userModel).Error
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}
func (r *UserRepository) Add(user *model.User) error {
	err := r.db.Table("users").Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) List() ([]*model.User, error) {
	var userModel []*model.User
	err := r.db.Table("users").Where("deleted_at IS NULL").Find(&userModel).Error
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func (r *UserRepository) ChangeStatus(request *contract.ChangeStatusRes) error {
	err := r.db.Table("users").Where("id = ?", request.Id).Update("status_user_id", request.StatusUserId).Error
	if err != nil {
		return err
	}
	return nil
}
