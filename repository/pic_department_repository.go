package repository

import (
	"errors"

	"ticket.narindo.com/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type picDepartmentRepository struct {
	repo *gorm.DB
}

type PicDepartmentRepository interface {
	Insert(dept *model.PicDepartment) error
	FindAllBy(by map[string]interface{}) ([]model.PicDepartment, error)
	FindAllWithPreload(by map[string]interface{}) ([]model.PicDepartment, error)
	UpdateByModel(payload *model.PicDepartment) error
	Delete(dept *model.PicDepartment) error
}

func (r *picDepartmentRepository) Insert(dept *model.PicDepartment) error {
	res := r.repo.Create(dept).Error
	return res
}

func (r *picDepartmentRepository) FindAllBy(by map[string]interface{}) ([]model.PicDepartment, error) {
	var picDepartments []model.PicDepartment
	res := r.repo.Where(by).Find(&picDepartments)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return picDepartments, nil
		} else {
			return picDepartments, err
		}
	}
	return picDepartments, nil
}

func (r *picDepartmentRepository) FindAllWithPreload(by map[string]interface{}) ([]model.PicDepartment, error) {
	var picDepartments []model.PicDepartment
	res := r.repo.Preload(clause.Associations).Where(by).Find(&picDepartments)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return picDepartments, nil
		} else {
			return picDepartments, err
		}
	}
	return picDepartments, nil
}

func (r *picDepartmentRepository) UpdateByModel(payload *model.PicDepartment) error {
	res := r.repo.Model(&payload).Updates(payload).Error
	return res
}

func (r *picDepartmentRepository) Delete(dept *model.PicDepartment) error {
	res := r.repo.Delete(&model.PicDepartment{}, dept).Error
	return res
}

func InitPicDepartmentRepository(db *gorm.DB) PicDepartmentRepository {
	picDepartmentRepo := new(picDepartmentRepository)
	picDepartmentRepo.repo = db
	return picDepartmentRepo
}
