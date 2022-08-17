package manager

import "ticket.narindo.com/repository"

type repositoryManager struct {
	infra Infra
}

type RepositoryManager interface {
	UserRepository() repository.UserRepository
	PicRepository() repository.PicRepository
	PicDepartmentRepository() repository.PicDepartmentRepository
	UserRoleRepository() repository.UserRoleRepository
	RoleRepository() repository.RoleRepository
}

func (r *repositoryManager) UserRepository() repository.UserRepository {
	return repository.InitUserRepository(r.infra.SqlDb())
}
func (r *repositoryManager) PicRepository() repository.PicRepository {
	return repository.InitPicRepository(r.infra.SqlDb())
}
func (r *repositoryManager) UserRoleRepository() repository.UserRoleRepository {
	return repository.InitUserRoleRepository(r.infra.SqlDb())
}
func (r *repositoryManager) RoleRepository() repository.RoleRepository {
	return repository.InitRoleRepository(r.infra.SqlDb())
}
func (r *repositoryManager) PicDepartmentRepository() repository.PicDepartmentRepository {
	return repository.InitPicDepartmentRepository(r.infra.SqlDb())
}

func InitRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{infra: infra}
}
