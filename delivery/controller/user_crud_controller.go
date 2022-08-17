package controller

import (
	"fmt"
	"net/http"

	"ticket.narindo.com/manager"
	"ticket.narindo.com/model"

	"github.com/gin-gonic/gin"
)

type UserCrudController struct {
	router         *gin.Engine
	usecaseManager manager.UsecaseManager
}

func (u *UserCrudController) getAllUsers(c *gin.Context) {
	result, err := u.usecaseManager.UserCrudUsecase().GetAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when retrieving users",
		})
		return
	}

	var formattedResult = func(input []model.User) []map[string]string {
		var fr []map[string]string
		for _, i := range input {
			if i.Pic.DepartmentId == 0 {
				fr = append(fr, map[string]string{"userId": i.UserId, "userEmail": i.Email, "userName": i.Name, "createdAt": i.CreatedAt, "isActive": fmt.Sprintf("%v", i.IsActive), "roleName": i.UserRole.Role.RoleName})
			} else if i.Pic.DepartmentId != 0 {
				fr = append(fr, map[string]string{"userId": i.UserId, "userEmail": i.Email, "userName": i.Name, "createdAt": i.CreatedAt, "isActive": fmt.Sprintf("%v", i.IsActive), "roleName": i.UserRole.Role.RoleName, "picDepartment": fmt.Sprintf("%v", i.Pic.PicDepartment.DepartmentName)})

			}
		}
		return fr
	}(result)
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": formattedResult,
		// "message": result,
	})
}

func (u *UserCrudController) getAllRoles(c *gin.Context) {
	result, err := u.usecaseManager.UserCrudUsecase().GetAllRoles()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when retrieving roles",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": result,
	})
}

func (u *UserCrudController) getAllDepartments(c *gin.Context) {
	result, err := u.usecaseManager.UserCrudUsecase().GetAllDepartments()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when retrieving departments",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": result,
	})
}

func (u *UserCrudController) insertNewUserAsCustomerService(c *gin.Context) {
	var newUser *model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		err := u.usecaseManager.UserCrudUsecase().InsertNewUser(newUser.Name, newUser.Email, "CUSTOMER SERVICE")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating new user as customer service",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"message": fmt.Sprintf("Created new user as CUSTOMER SERVICE with email: %s", newUser.Email),
		})
	}
}

func (u *UserCrudController) insertNewUserAsSuperAdmin(c *gin.Context) {
	var newUser *model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		err := u.usecaseManager.UserCrudUsecase().InsertNewUser(newUser.Name, newUser.Email, "SUPER ADMIN")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating new user as super admin",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"message": fmt.Sprintf("Created new user as SUPER ADMIN with email: %s", newUser.Email),
		})
	}
}
func (u *UserCrudController) insertNewUserAsPic(c *gin.Context) {
	var newUser *model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		err := u.usecaseManager.UserCrudUsecase().InsertNewUserAsPic(newUser.Name, newUser.Email, "PIC", newUser.Pic.DepartmentId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating new user as super admin",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"message": fmt.Sprintf("Created new user as PIC with email: %s", newUser.Email),
		})
	}
}

func (u *UserCrudController) changeUserActiveStatus(c *gin.Context) {
	userId := c.Param("user-id")
	err := u.usecaseManager.UserCrudUsecase().ChangeUserActiveStatus(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when changing user status",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": fmt.Sprintf("%s status changed", userId),
	})
}

func (u *UserCrudController) changeUserEmail(c *gin.Context) {
	userId := c.Param("user-id")
	var user *model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		err := u.usecaseManager.UserCrudUsecase().UpdateUserEmail(userId, user.Email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "SUCCESS",
				"message": fmt.Sprintf("Changed user %s email to %s", userId, user.Email),
			})
		}
	}
}

func (u *UserCrudController) changeUserName(c *gin.Context) {
	userId := c.Param("user-id")
	var user *model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		err := u.usecaseManager.UserCrudUsecase().UpdateUserName(userId, user.Name)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "SUCCESS",
				"message": fmt.Sprintf("Changed user %s name to %s", userId, user.Name),
			})
		}
	}
}

func (u *UserCrudController) changeUserRole(c *gin.Context) {
	userId := c.Param("user-id")
	var newRole *model.Role
	if err := c.BindJSON(&newRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		err := u.usecaseManager.UserCrudUsecase().UpdateUserRole(userId, newRole.RoleName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "SUCCESS",
				"message": fmt.Sprintf("Changed user %s role to %s", userId, newRole.RoleName),
			})
		}
	}
}
func (u *UserCrudController) changeUserPicDepartment(c *gin.Context) {
	userId := c.Param("user-id")
	var newDep *model.PicDepartment
	if err := c.BindJSON(&newDep); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		err := u.usecaseManager.UserCrudUsecase().UpdateUserPicDepartment(userId, newDep.DepartmentName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "SUCCESS",
				"message": fmt.Sprintf("Changed pic %s department to %s", userId, newDep.DepartmentName),
			})
		}
	}
}
func InitUserCrudController(router *gin.Engine, um manager.UsecaseManager) *UserCrudController {
	controller := UserCrudController{
		router:         router,
		usecaseManager: um,
	}
	userRouter := router.Group("api/user")
	userRouter.GET("/list", controller.getAllUsers)
	userRouter.POST("/create/cs", controller.insertNewUserAsCustomerService)
	userRouter.POST("/create/sa", controller.insertNewUserAsSuperAdmin)
	userRouter.POST("/create/pic", controller.insertNewUserAsPic)
	userRouter.PUT("/status/update/:user-id", controller.changeUserActiveStatus)
	userRouter.PUT("/email/update/:user-id", controller.changeUserEmail)
	userRouter.PUT("/name/update/:user-id", controller.changeUserName)
	userRouter.GET("/role", controller.getAllRoles)
	userRouter.PUT("/role/update/:user-id", controller.changeUserRole)
	userRouter.GET("/department/", controller.getAllDepartments)
	userRouter.PUT("/department/update/:user-id", controller.changeUserPicDepartment)
	return &controller
}
