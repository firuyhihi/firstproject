package usecase

import (
	"errors"
	"fmt"

	"ticket.narindo.com/model"
	"ticket.narindo.com/repository"
	"ticket.narindo.com/utils"
)

type TicketUseCase interface {
	CreateTicket(ticket *model.Ticket) error
	ListByUser(userRole *model.UserRole) ([]model.Ticket, error)
	ListByDepartment(departmentId int) ([]model.Ticket, error)

	GetById(ticketId string) (model.Ticket, error)

	UpdatePIC(ticketId string, picId int) error
	UpdateStatus(ticketId string, statusId int) error

	GetTicketStatus(ticketId string) (model.Ticket, error)
	GetTicketList() ([]model.Ticket, error)
	GetNotificationList() ([]model.Ticket, error)
	GetTicketSummary() ([]model.Ticket, error)

	// ListByDate(userId, orderBy string) ([]model.Ticket, error)
	// ListByPriority(userId, priority string) ([]model.Ticket, error)
	// ListByCategory(userId, category string) ([]model.Ticket, error)

	// FilterByDate(userId string, startDate, endDate time.Time) ([]model.Ticket, error)
	// GetTotalOpen(userId string) (int64, error)
	// GetTotalClose(userId string) (int64, error)
	// GetTotalProgress(userId string) (int64, error)
}

type ticketUseCase struct {
	repo     repository.TicketRepository
	repoDept repository.PicDepartmentRepository
	repoPic  repository.PicRepository
}

func (t *ticketUseCase) ListByDepartment(departmentId int) ([]model.Ticket, error) {
	ticketList, err := t.repo.FindAllBy(map[string]interface{}{"department_id": departmentId, "status_id": 1})
	if err != nil {
		return []model.Ticket{}, err
	}
	return ticketList, nil
}

func (t *ticketUseCase) CreateTicket(ticket *model.Ticket) error {
	var tocheck = make(map[string]string)
	tocheck["TicketSubject"] = ticket.TicketSubject
	tocheck["TicketMessage"] = ticket.TicketMessage
	tocheck["CsId"] = ticket.CsId

	rescheck := utils.CheckEmptyStringRequest(tocheck)
	if rescheck != "" {
		return fmt.Errorf("create failed, missing required field %v", rescheck)
	}

	if ticket.CsId == "" || ticket.DepartmentId == 0 || ticket.TicketSubject == "" || ticket.TicketMessage == "" {
		return errors.New("create failed, missing required field")
	}
	// status, cs, priority
	department, err := t.repoDept.FindAllBy(map[string]interface{}{"id": ticket.DepartmentId})
	if err != nil {
		return errors.New("problem exists")
	}
	if len(department) == 0 {
		return errors.New("create failed, departement not found")
	} //tambahin cek

	status, err := t.repo.FindAllBy(map[string]interface{}{"id": ticket.StatusId})
	if err != nil {
		return errors.New("problem exists")
	}
	if len(status) == 0 {
		return errors.New("create failed, status not found")
	}

	csId, err := t.repo.FindAllBy(map[string]interface{}{"id": ticket.CsId})
	if err != nil {
		return errors.New("problem exists")
	}
	if len(csId) == 0 {
		return errors.New("create failed, cs id not found")
	}

	priority, err := t.repo.FindAllBy(map[string]interface{}{"id": ticket.PriorityId})
	if err != nil {
		return errors.New("problem exists")
	}
	if len(priority) == 0 {
		return errors.New("create failed, priority not found")
	}

	var newTicketId = utils.GenerateId(department[0].DepartmentName)
	ticket.TicketId = newTicketId

	ticket.StatusId = 1
	ticket.IsActive = true

	err = t.repo.Create(ticket)

	return err
}

func (t *ticketUseCase) ListByUser(userRole *model.UserRole) ([]model.Ticket, error) {
	if userRole.RoleId == 2 {
		ticketList, err := t.repo.FindAllBy(map[string]interface{}{"cs_id": userRole.UserId})
		if err != nil {
			return []model.Ticket{}, err
		}
		return ticketList, nil
	}

	pic, _ := t.repoPic.FindAllBy(map[string]interface{}{"user_id": userRole.UserId})
	if len(pic) == 0 {
		return []model.Ticket{}, errors.New("can not find pic")
	}

	ticketList, err := t.repo.FindAllBy(map[string]interface{}{"pic_id": pic[0].Id})
	if err != nil {
		return []model.Ticket{}, err
	}
	return ticketList, nil

}

// func (t *ticketUseCase) ListByDate(userId, orderBy string) ([]model.Ticket, error) {
// 	return t.repo.ListByDate(userId, orderBy)
// }

// func (t *ticketUseCase) ListByPriority(userId, priority string) ([]model.Ticket, error) {
// 	return t.repo.ListByPriority(userId, priority)
// }

// func (t *ticketUseCase) ListByCategory(userId, category string) ([]model.Ticket, error) {
// 	return t.repo.ListByPriority(userId, category)
// }

// func (t *ticketUseCase) FilterByDate(userId string, startDate, endDate time.Time) ([]model.Ticket, error) {
// 	return t.repo.FilterByDate(userId, startDate, endDate)
// }

func (t *ticketUseCase) GetById(ticketId string) (model.Ticket, error) {
	return t.repo.FindBy(map[string]interface{}{"ticket_id": ticketId})
}

func (t *ticketUseCase) GetTicketStatus(ticketId string) (model.Ticket, error) {
	return t.repo.FindBy(map[string]interface{}{"ticket_id": ticketId})
}

func (t *ticketUseCase) GetTicketList() ([]model.Ticket, error) {
	return t.repo.FindAllBy(map[string]interface{}{"ticket_id": ""})
}

func (t *ticketUseCase) GetNotificationList() ([]model.Ticket, error) {
	return t.repo.FindAllBy(map[string]interface{}{"ticket_id": ""})
}

func (t *ticketUseCase) GetTicketSummary() ([]model.Ticket, error) {
	return t.repo.FindAllBy(map[string]interface{}{"ticket_id": ""})
}

// func (t *ticketUseCase) GetTotalOpen(ticketId string) (int64, error) {
// 	return t.repo.GetTotalOpen(ticketId)
// }

// func (t *ticketUseCase) GetTotalClose(ticketId string) (int64, error) {
// 	return t.repo.GetTotalClose(ticketId)
// }

// func (t *ticketUseCase) GetTotalProgress(ticketId string) (int64, error) {
// 	return t.repo.GetTotalProgress(ticketId)
// }

func (t *ticketUseCase) UpdatePIC(ticketId string, picId int) error {
	return t.repo.UpdateBy(map[string]interface{}{"ticket_id": ticketId}, map[string]interface{}{"pic_id": picId})
}

func (t *ticketUseCase) UpdateStatus(ticketId string, statusId int) error {
	return t.repo.UpdateBy(map[string]interface{}{"ticket_id": ticketId}, map[string]interface{}{"status_id": statusId})
}

func NewTicketUseCase(repo repository.TicketRepository, repoDept repository.PicDepartmentRepository, repoPic repository.PicRepository) TicketUseCase {
	return &ticketUseCase{
		repo:     repo,
		repoDept: repoDept,
		repoPic:  repoPic,
	}
}
