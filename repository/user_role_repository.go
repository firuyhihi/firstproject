package repository

import (
	"errors"

	"ticket.narindo.com/model"

	"gorm.io/gorm"
)

type userRoleRepository struct {
	repo *gorm.DB
}

type UserRoleRepository interface {
	Insert(userRole *model.UserRole) error
	FindAllBy(by map[string]interface{}) ([]model.UserRole, error)
	FindAllWithPreload(by map[string]interface{}, preload string) ([]model.UserRole, error)
	UpdateBy(by map[string]interface{}, value map[string]interface{}) error
}

func (u *userRoleRepository) Insert(userRole *model.UserRole) error {
	return u.repo.Create(userRole).Error
}

func (u *userRoleRepository) FindAllBy(by map[string]interface{}) ([]model.UserRole, error) {
	var userRoles []model.UserRole
	res := u.repo.Where(by).Find(&userRoles)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userRoles, nil
		} else {
			return userRoles, err
		}
	}
	return userRoles, nil
}

func (u *userRoleRepository) FindAllWithPreload(by map[string]interface{}, preload string) ([]model.UserRole, error) {
	var userRoles []model.UserRole
	res := u.repo.Preload(preload).Where(by).Find(&userRoles)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userRoles, nil
		} else {
			return userRoles, err
		}
	}
	return userRoles, nil
}

func (u *userRoleRepository) UpdateBy(by map[string]interface{}, value map[string]interface{}) error {
	return u.repo.Model(model.UserRole{}).Where(by).Updates(value).Error
}

// func (u *userRoleRepository) Delete(userRole *model.UserRole) error  {
// 	return u.repo.Create(userRole).Error
// }

func InitUserRoleRepository(db *gorm.DB) UserRoleRepository {
	userRoleRepo := new(userRoleRepository)
	userRoleRepo.repo = db
	return userRoleRepo
}
