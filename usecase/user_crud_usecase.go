package usecase

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"ticket.narindo.com/model"
	"ticket.narindo.com/repository"
)

type userCrudUsecase struct {
	userRepo     repository.UserRepository
	userRoleRepo repository.UserRoleRepository
	roleRepo     repository.RoleRepository
	picRepo      repository.PicRepository
	picDepRepo   repository.PicDepartmentRepository
}

type UserCrudUsecase interface {
	GetAllUsers() ([]model.User, error)
	GetAllRoles() ([]model.Role, error)
	GetAllDepartments() ([]model.PicDepartment, error)
	InsertNewUser(name, email, role string) error
	InsertNewUserAsPic(name, email, role string, department int) error
	ChangeUserActiveStatus(email string) error
	UpdateUserEmail(userId string, newEmail string) error
	UpdateUserName(userId string, newName string) error
	UpdateUserRole(userId string, newRole string) error
	UpdateUserPicDepartment(userId string, newDepartment string) error
}

func (u *userCrudUsecase) GetAllUsers() ([]model.User, error) {
	return u.userRepo.FindAllWithPreload(map[string]interface{}{})
}

func (u *userCrudUsecase) GetAllRoles() ([]model.Role, error) {
	return u.roleRepo.FindAllWithPreload(map[string]interface{}{})
}

func (u *userCrudUsecase) GetAllDepartments() ([]model.PicDepartment, error) {
	return u.picDepRepo.FindAllWithPreload(map[string]interface{}{})
}

func (u *userCrudUsecase) InsertNewUser(name, email, role string) error {
	users, err := u.userRepo.FindAllBy(map[string]interface{}{})
	if err != nil {
		return err
	} else if err == nil {
		for _, u := range users {
			if u.Email == email {
				return errors.New("can not create a new user with same email")
			}
		}
	}

	newId := userIdGenerator()
	var inputRole model.Role
	roles, err := u.roleRepo.FindAllBy(map[string]interface{}{})
	for _, v := range roles {
		if v.RoleName == role {
			inputRole = v
			break
		}
	}

	newUserRole := model.UserRole{UserId: newId, RoleId: inputRole.Id}
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.UserId == newId {
			newId = userIdGenerator()
		}
	}
	newUser := model.User{
		UserId:    newId,
		Email:     email,
		Name:      name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		IsActive:  true,
		UserRole:  newUserRole,
	}
	err = u.userRepo.Insert(&newUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCrudUsecase) InsertNewUserAsPic(name, email, role string, department int) error {
	users, err := u.userRepo.FindAllBy(map[string]interface{}{})
	if err != nil {
		return err
	} else if err == nil {
		for _, u := range users {
			if u.Email == email {
				return errors.New("can not create a new user with same email")
			}
		}
	}

	newId := userIdGenerator()
	var inputRole model.Role
	var inputDep model.PicDepartment
	roles, err := u.roleRepo.FindAllBy(map[string]interface{}{})
	if err != nil {
		return err
	}
	for _, v := range roles {
		if v.RoleName == role {
			inputRole = v
			break
		}
	}
	dept, err := u.picDepRepo.FindAllBy(map[string]interface{}{})
	if err != nil {
		return err
	}
	for _, v := range dept {
		if v.Id == department {
			inputDep = v
			break
		}
	}

	newUserRole := model.UserRole{UserId: newId, RoleId: inputRole.Id}
	newPic := model.Pic{RoleId: inputRole.Id, UserId: newId, DepartmentId: inputDep.Id}

	for _, u := range users {
		if u.UserId == newId {
			newId = userIdGenerator()
		}
	}
	newUser := model.User{
		UserId:    newId,
		Email:     email,
		Name:      name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		IsActive:  true,
		UserRole:  newUserRole,
		Pic:       newPic,
	}
	err = u.userRepo.Insert(&newUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCrudUsecase) ChangeUserActiveStatus(userId string) error {
	user, err := u.userRepo.FindBy(map[string]interface{}{"user_id": userId})
	if err != nil {
		return err
	}
	switch user.IsActive {
	case true:
		return u.userRepo.UpdateBy(map[string]interface{}{"user_id": userId}, map[string]interface{}{"is_active": false})

	case false:
		return u.userRepo.UpdateBy(map[string]interface{}{"user_id": userId}, map[string]interface{}{"is_active": true})
	}

	return nil
}

func (u *userCrudUsecase) UpdateUserEmail(userId string, newEmail string) error {
	user, err := u.userRepo.FindBy(map[string]interface{}{"user_id": userId})
	if err != nil {
		return err
	}
	return u.userRepo.UpdateBy(map[string]interface{}{"user_id": user.UserId}, map[string]interface{}{"email": newEmail})
}

func (u *userCrudUsecase) UpdateUserName(userId string, newName string) error {
	user, err := u.userRepo.FindBy(map[string]interface{}{"user_id": userId})
	if err != nil {
		return err
	}
	return u.userRepo.UpdateBy(map[string]interface{}{"user_id": user.UserId}, map[string]interface{}{"name": newName})
}

func (u *userCrudUsecase) UpdateUserRole(userId string, newRole string) error {
	user, err := u.userRepo.FindBy(map[string]interface{}{"user_id": userId})
	if err != nil {
		return err
	}

	var inputRole model.Role
	roles, err := u.roleRepo.FindAllBy(map[string]interface{}{})
	if err != nil {
		return err
	}
	for _, v := range roles {
		if v.RoleName == newRole {
			inputRole = v
			break
		}
	}
	return u.userRoleRepo.UpdateBy(map[string]interface{}{"user_id": user.UserId}, map[string]interface{}{"role_id": inputRole.Id})
}

func (u *userCrudUsecase) UpdateUserPicDepartment(userId string, newDepartment string) error {
	user, err := u.userRepo.FindBy(map[string]interface{}{"user_id": userId})
	if err != nil {
		return err
	}

	var inputDep model.PicDepartment
	deps, err := u.picDepRepo.FindAllBy(map[string]interface{}{})
	if err != nil {
		return err
	}
	for _, v := range deps {
		if v.DepartmentName == newDepartment {
			inputDep = v
			break
		}
	}
	return u.picRepo.UpdateBy(map[string]interface{}{"user_id": user.UserId}, map[string]interface{}{"department_id": inputDep.Id})
}

func InitUserUsecase(u repository.UserRepository, ur repository.UserRoleRepository, r repository.RoleRepository, p repository.PicRepository, pd repository.PicDepartmentRepository) UserCrudUsecase {
	userUsec := new(userCrudUsecase)
	userUsec.userRepo = u
	userUsec.userRoleRepo = ur
	userUsec.roleRepo = r
	userUsec.picRepo = p
	userUsec.picDepRepo = pd
	return userUsec
}

func userIdGenerator() string {
	var v [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		v[i] = rand.Intn(9)
	}
	res := 0
	op := 1
	for i := len(v) - 1; i >= 0; i-- {
		res += v[i] * op
		op *= 10
	}
	newId := strconv.Itoa(res)
	return newId
}
