package repository

import (
	"errors"

	"ticket.narindo.com/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type roleRepository struct {
	repo *gorm.DB
}

type RoleRepository interface {
	Insert(role *model.Role) error
	FindAllBy(by map[string]interface{}) ([]model.Role, error)
	FindAllWithPreload(by map[string]interface{}) ([]model.Role, error)
	UpdateByModel(payload *model.Role) error
	Delete(role *model.Role) error
}

func (r *roleRepository) Insert(role *model.Role) error {
	res := r.repo.Create(role).Error
	return res
}

func (r *roleRepository) FindAllBy(by map[string]interface{}) ([]model.Role, error) {
	var roles []model.Role
	res := r.repo.Where(by).Find(&roles)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return roles, nil
		} else {
			return roles, err
		}
	}
	return roles, nil
}

func (r *roleRepository) FindAllWithPreload(by map[string]interface{}) ([]model.Role, error) {
	var roles []model.Role
	res := r.repo.Preload(clause.Associations).Where(by).Find(&roles)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return roles, nil
		} else {
			return roles, err
		}
	}
	return roles, nil
}

func (r *roleRepository) UpdateByModel(payload *model.Role) error {
	res := r.repo.Model(&payload).Updates(payload).Error
	return res
}

func (r *roleRepository) Delete(role *model.Role) error {
	res := r.repo.Delete(&model.Role{}, role).Error
	return res
}

func InitRoleRepository(db *gorm.DB) RoleRepository {
	roleRepo := new(roleRepository)
	roleRepo.repo = db
	return roleRepo
}
