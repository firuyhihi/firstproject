package controller

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"gorm.io/gorm"
	"ticket.narindo.com/delivery/api"

	"ticket.narindo.com/model"
	"ticket.narindo.com/usecase"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	router   *gin.Engine
	ucTicket usecase.TicketUseCase
	api.BaseApi
}

func (t *TicketController) createTicket(c *gin.Context) {
	var newTicket model.Ticket
	err := t.ParseRequestBody(c, &newTicket)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err,
		})
		return
	}
	err = t.ucTicket.CreateTicket(&newTicket)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": newTicket,
	})
}

func (t *TicketController) ListTicketByUserId(c *gin.Context) {
	var userRole model.UserRole
	err := t.ParseRequestBody(c, &userRole)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err,
		})
		return
	}

	print, err := t.ucTicket.ListByUser(&userRole)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": print,
	})
}

func (t *TicketController) ListTicketByDepartmentId(c *gin.Context) {
	departmentId := c.Param("id")
	integerDeptId, _ := strconv.Atoi(departmentId)
	print, err := t.ucTicket.ListByDepartment(integerDeptId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": print,
	})
}

// func (t *TicketController) listTicketByDate(c *gin.Context) {
// 	userId := ""
// 	orderBy := c.Param("orderBy")

// 	print, err := t.ucTicket.ListByDate(userId, orderBy)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			t.Failed(c, utils.DataNotFoundError())
// 			return
// 		}
// 		t.Failed(c, err)
// 		return
// 	}
// 	t.Success(c, print)
// }

// func (t *TicketController) filterTicketByDate(c *gin.Context) {
// 	userId := ""
// 	var filterDate struct {
// 		// startDate string
// 		startTime string
// 		// endDate string
// 		endTime string
// 	}

// 	if err := c.BindJSON(&filterDate); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  "BAD REQUEST",
// 			"message": err.Error(),
// 		})
// 	}

// 	// startTime := c.Param("startTime") + "2006-01-02 00:00:00"

// 	startDate, err := time.Parse("2006-01-02 15:04:05", filterDate.startTime)
// 	if err != nil {
// 		t.Failed(c, err)
// 		return
// 	}

// 	// endTime := c.Param("endTime") + "2006-01-02 00:00:00"

// 	endDate, err := time.Parse("2006-01-02 15:04:05", filterDate.endTime)
// 	if err != nil {
// 		t.Failed(c, err)
// 		return
// 	}

// 	print, err := t.ucTicket.FilterByDate(userId, startDate, endDate)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			t.Failed(c, utils.DataNotFoundError())
// 			return
// 		}
// 		t.Failed(c, err)
// 		return
// 	}
// 	t.Success(c, print)
// }

// func (t *TicketController) listTicketByPriority(c *gin.Context) {
// 	// nanti diambil dari jwt aja bisa
// 	// userId := ""
// 	userId := c.Param("userId")
// 	print, err := t.ucTicket.ListByUser(userId)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			t.Failed(c, utils.DataNotFoundError())
// 			return
// 		}
// 		t.Failed(c, err)
// 		return
// 	}
// 	t.Success(c, print)
// }

// func (t *TicketController) listTicketByCategory(c *gin.Context) {
// 	// nanti diambil dari jwt aja bisa
// 	// userId := ""
// 	userId := c.Param("userId")
// 	print, err := t.ucTicket.ListByUser(userId)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			t.Failed(c, utils.DataNotFoundError())
// 			return
// 		}
// 		t.Failed(c, err)
// 		return
// 	}
// 	t.Success(c, print)
// }

func (t *TicketController) getTicketById(c *gin.Context) {
	ticketId := c.Param("id")
	print, err := t.ucTicket.GetById(ticketId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error ticket not found",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": print,
	})
}

// func (t *TicketController) getTotalOpenTicket(c *gin.Context) {
// 	userId := ""
// 	print, err := t.ucTicket.GetTotalOpen(userId)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			t.Failed(c, utils.DataNotFoundError())
// 			return
// 		}
// 		t.Failed(c, err)
// 		return
// 	}
// 	t.Success(c, print)
// }

// func (t *TicketController) getTotalCloseTicket(c *gin.Context) {
// 	userId := ""
// 	print, err := t.ucTicket.GetTotalClose(userId)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			t.Failed(c, utils.DataNotFoundError())
// 			return
// 		}
// 		t.Failed(c, err)
// 		return
// 	}
// 	t.Success(c, print)
// }

// func (t *TicketController) getTotalOnProgressTicket(c *gin.Context) {
// 	userId := ""
// 	print, err := t.ucTicket.GetTotalProgress(userId)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			t.Failed(c, utils.DataNotFoundError())
// 			return
// 		}
// 		t.Failed(c, err)
// 		return
// 	}
// 	t.Success(c, print)
// }

func (t *TicketController) updatePIC(c *gin.Context) {
	var updateInput struct {
		TicketId string `json:"ticketId"`
		PicId    int    `json:"picId"`
	}

	if err := c.BindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		fmt.Println(updateInput)
		err := t.ucTicket.UpdatePIC(updateInput.TicketId, updateInput.PicId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when assign new PIC",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"message": "complain assigned to a new pic",
		})
	}
}

func (t *TicketController) updateStatus(c *gin.Context) {
	var updateInput struct {
		TicketId string `json:"ticketId"`
		StatusId    int    `json:"statusId"`
	}

	if err := c.BindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		fmt.Println(updateInput)
		err := t.ucTicket.UpdateStatus(updateInput.TicketId, updateInput.StatusId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when update status ticket",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"message": "updated status success",
		})
	}
}

func NewTicketController(router *gin.Engine, ucTicket usecase.TicketUseCase) *TicketController {
	controller := TicketController{
		router:   router,
		ucTicket: ucTicket,
	}

	// router.PUT("/updatePIC/:userId", controller.updatePIC)

	protectedGroup := router.Group("api/ticket")

	protectedGroup.POST("", controller.createTicket)
	protectedGroup.GET("/list", controller.ListTicketByUserId)
	protectedGroup.GET("/department/:id", controller.ListTicketByDepartmentId)

	// protectedGroup.GET("/listByDate/:orderBy", controller.listTicketByDate)
	// protectedGroup.GET("/ticketListByDate/:dateTime", controller.tiketListByDate)
	// protectedGroup.GET("/listByPriority/:priority", controller.listTicketByPriority)
	// protectedGroup.GET("/listByCategory/:category", controller.listTicketByCategory)

	protectedGroup.PUT("/update-pic", controller.updatePIC)
	protectedGroup.PUT("/update-status", controller.updateStatus)

	// protectedGroup.GET("/filterByDate/:startDate/:endDate", controller.filterTicketByDate)

	protectedGroup.GET("/:id", controller.getTicketById)
	// protectedGroup.GET("/getTotalOpen", controller.getTotalOpenTicket)
	// protectedGroup.GET("/getTotalClose", controller.getTotalCloseTicket)
	// protectedGroup.GET("/getTotalProgress", controller.getTotalOnProgressTicket)
	return &controller
}
