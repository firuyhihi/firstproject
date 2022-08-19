package manager

import "ticket.narindo.com/usecase"

type usecaseManager struct {
	repoManager RepositoryManager
}

type UsecaseManager interface {
	UserCrudUsecase() usecase.UserCrudUsecase
	// LoginUsecase() usecase.LoginUsecase
	TicketUseCase() usecase.TicketUseCase
}

func (u *usecaseManager) UserCrudUsecase() usecase.UserCrudUsecase {
	return usecase.InitUserUsecase(u.repoManager.UserRepository(), u.repoManager.UserRoleRepository(), u.repoManager.RoleRepository(), u.repoManager.PicRepository(), u.repoManager.PicDepartmentRepository())

}

func (u *usecaseManager) TicketUseCase() usecase.TicketUseCase {
	return usecase.NewTicketUseCase(u.repoManager.TicketRepo(), u.repoManager.PicDepartmentRepository(), u.repoManager.PicRepository())
}

// func (u *usecaseManager) LoginUsecase() usecase.LoginUsecase {
// 	return usecase.InitLoginUsecase(u.repoManager.UserRepository())

// }

func InitUsecasesManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
