package usecase

import (
	"ticket.narindo.com/model"
	"ticket.narindo.com/repository"
)

type picUsecase struct {
	repo repository.PicRepository
}

type PicUsecase interface {
	GetAllPic() ([]model.Pic, error)
}

func (u *picUsecase) GetAllPic() ([]model.Pic, error) {
	return u.repo.FindAllWithPreload(map[string]interface{}{}, "User", "UserRole", "PicDepartment")
}

func InitPicUsecases(p repository.PicRepository) PicUsecase {
	picUsec := new(picUsecase)
	picUsec.repo = p
	return picUsec
}
