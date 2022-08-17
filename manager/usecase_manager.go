package manager

import "ticket.narindo.com/usecase"

type usecaseManager struct {
	repoManager RepositoryManager
}

type UsecaseManager interface {
	UserCrudUsecase() usecase.UserCrudUsecase
	// LoginUsecase() usecase.LoginUsecase
}

func (u *usecaseManager) UserCrudUsecase() usecase.UserCrudUsecase {
	return usecase.InitUserUsecase(u.repoManager.UserRepository(), u.repoManager.UserRoleRepository(), u.repoManager.RoleRepository(), u.repoManager.PicRepository(), u.repoManager.PicDepartmentRepository())

}

// func (u *usecaseManager) LoginUsecase() usecase.LoginUsecase {
// 	return usecase.InitLoginUsecase(u.repoManager.UserRepository())

// }

func InitUsecasesManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
