package repository

import (
	"errors"

	"ticket.narindo.com/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	repo *gorm.DB
}

type UserRepository interface {
	Insert(user *model.User) error
	FindBy(by map[string]interface{}) (model.User, error)
	FindWithPreload(by map[string]interface{}) (model.User, error)
	FindAllBy(by map[string]interface{}) ([]model.User, error)
	FindAllWithPreload(by map[string]interface{}) ([]model.User, error)
	UpdateBy(by map[string]interface{}, value map[string]interface{}) error
	Delete(user *model.User) error
}

func (t *userRepository) Insert(user *model.User) error {
	res := t.repo.Create(user).Error
	return res
}

func (u *userRepository) FindBy(by map[string]interface{}) (model.User, error) {
	var user model.User
	res := u.repo.Where(by).First(&user)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		} else {
			return user, err
		}
	}
	return user, nil
}

func (u *userRepository) FindWithPreload(by map[string]interface{}) (model.User, error) {
	var user model.User
	res := u.repo.Preload("UserRole.Role").Preload(clause.Associations).First(&user)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		} else {
			return user, err
		}
	}
	return user, nil
}

func (t *userRepository) FindAllBy(by map[string]interface{}) ([]model.User, error) {
	var users []model.User
	res := t.repo.Where(by).Find(&users)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		} else {
			return users, err
		}
	}
	return users, nil
}

func (u *userRepository) FindAllWithPreload(by map[string]interface{}) ([]model.User, error) {
	var users []model.User
	res := u.repo.Preload("UserRole.Role").Preload("Pic.PicDepartment").Preload(clause.Associations).Where(by).Find(&users)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		} else {
			return users, err
		}
	}
	return users, nil
}

func (u *userRepository) UpdateBy(by map[string]interface{}, value map[string]interface{}) error {
	return u.repo.Model(model.User{}).Where(by).Updates(value).Error
}

func (u *userRepository) Delete(user *model.User) error {
	res := u.repo.Delete(&model.User{}, user).Error
	return res
}

func InitUserRepository(db *gorm.DB) UserRepository {
	userRepo := new(userRepository)
	userRepo.repo = db
	return userRepo
}
