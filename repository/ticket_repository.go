package repository

import (
	"errors"

	"ticket.narindo.com/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TicketRepository interface {
	Create(newTicket *model.Ticket) error
	// ListByUser(user string) ([]model.Ticket, error)
	// ListByDate(userId, orderBy string) ([]model.Ticket, error)
	// ListByCategory(userId, category string) ([]model.Ticket, error)
	// ListByPriority(userId, priority string) ([]model.Ticket, error)

	// FilterByDate(userId string, startDate, endDate time.Time) ([]model.Ticket, error)

	// GetById(ticketId string) (model.Ticket, error)
	// GetTotalOpen(ticketId string) (int64, error)
	// GetTotalClose(ticketId string) (int64, error)
	// GetTotalProgress(ticketId string) (int64, error)

	// UpdatePIC(ticket *model.JsonTicket, userId string) error

	FindBy(by map[string]interface{}) (model.Ticket, error)
	FindAllBy(by map[string]interface{}) ([]model.Ticket, error)
	FindAll() ([]model.Ticket, error)
	UpdateBy(by map[string]interface{}, value map[string]interface{}) error
	Delete(ticket *model.Ticket) error
}

type ticketRepository struct {
	db *gorm.DB
}

func (t *ticketRepository) Delete(ticket *model.Ticket) error {
	res := t.db.Delete(&model.Ticket{}, ticket).Error
	return res
}

func (t *ticketRepository) FindAllBy(by map[string]interface{}) ([]model.Ticket, error) {
	var ticket []model.Ticket
	res := t.db.Preload(clause.Associations).Where(by).Find(&ticket)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ticket, nil
		} else {
			return ticket, err
		}
	}
	return ticket, nil
}

func (t *ticketRepository) FindAll() ([]model.Ticket, error) {
	var ticket []model.Ticket
	res := t.db.Preload(clause.Associations).Find(&ticket)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ticket, nil
		} else {
			return ticket, err
		}
	}
	return ticket, nil
}

func (t *ticketRepository) FindBy(by map[string]interface{}) (model.Ticket, error) {
	var ticket model.Ticket
	res := t.db.Preload(clause.Associations).Where(by).First(&ticket)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ticket, nil
		} else {
			return ticket, err
		}
	}
	return ticket, nil
}

func (t *ticketRepository) UpdateBy(by map[string]interface{}, value map[string]interface{}) error {
	return t.db.Model(model.Ticket{}).Where(by).Updates(value).Error
}

func (m *ticketRepository) Create(newTicket *model.Ticket) error {
	result := m.db.Create(newTicket).Error
	return result
}

// func (t *ticketRepository) ListByUser(userId string) ([]model.Ticket, error) {
// 	var ticket []model.Ticket
// 	err := t.db.Find(&ticket, "ticket_pic=?", userId).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return []model.Ticket{}, err
// 		}
// 		fmt.Println(err.Error())
// 		return []model.Ticket{}, fmt.Errorf("komplain dari user %v tidak ditemukan", userId)
// 	}
// 	return ticket, nil
// }

// func (t *ticketRepository) ListByDate(userId, orderBy string) ([]model.Ticket, error) {
// 	var ticket []model.Ticket
// 	err := t.db.Find(&ticket, "ticket_pic=?", userId).Order("ticket_date " + orderBy).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return []model.Ticket{}, err
// 		}
// 		fmt.Println(err.Error())
// 		return []model.Ticket{}, fmt.Errorf("komplain dari user %v tidak ditemukan", userId)
// 	}
// 	return ticket, nil
// }

// func (t *ticketRepository) ListByCategory(userId, category string) ([]model.Ticket, error) {
// 	var ticket []model.Ticket
// 	err := t.db.Find(&ticket).Where("ticket_pic=? and ticket_id like ('"+"%"+category+"')", userId).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return []model.Ticket{}, err
// 		}
// 		fmt.Println(err.Error())
// 		return []model.Ticket{}, fmt.Errorf("komplain dari user %v tidak ditemukan", userId)
// 	}
// 	return ticket, nil
// }

// func (t *ticketRepository) ListByPriority(userId, priority string) ([]model.Ticket, error) {
// 	var ticket []model.Ticket
// 	err := t.db.Find(&ticket).Where("ticket_pic=? and ticket_priority = ?", userId, priority).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return []model.Ticket{}, err
// 		}
// 		fmt.Println(err.Error())
// 		return []model.Ticket{}, fmt.Errorf("komplain dari user %v tidak ditemukan", userId)
// 	}
// 	return ticket, nil
// }

// func (t *ticketRepository) FilterByDate(userId string, startDate, endDate time.Time) ([]model.Ticket, error) {
// 	var ticket []model.Ticket
// 	err := t.db.Find(&ticket).Where("ticket_pic=? and ticket_date >= ? and ticket_date <= ?", userId, startDate, endDate).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return []model.Ticket{}, err
// 		}
// 		fmt.Println(err.Error())
// 		return []model.Ticket{}, fmt.Errorf("komplain dari user %v tidak ditemukan", userId)
// 	}
// 	return ticket, nil
// }

// func (t *ticketRepository) GetById(ticketId string) (model.Ticket, error) {
// 	var ticket model.Ticket
// 	err := t.db.Find(&ticket).Where("ticket_id = ?", ticketId).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return model.Ticket{}, err
// 		}
// 		fmt.Println(err.Error())
// 		return model.Ticket{}, fmt.Errorf("komplain dengan id %v tidak ditemukan", ticketId)
// 	}
// 	return ticket, nil
// }

// func (t *ticketRepository) GetTotalOpen(ticketId string) (int64, error) {
// 	var ticket model.Ticket
// 	var count int64
// 	// nanti diganti ticket_prioritynya
// 	err := t.db.Find(&ticket).Where("ticket_id = ? and ticket_priority = 'open'", ticketId).Count(&count).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return 0, err
// 		}
// 		fmt.Println(err.Error())
// 		return 0, fmt.Errorf("komplain dengan id %v tidak ditemukan", ticketId)
// 	}
// 	return count, nil
// }

// func (t *ticketRepository) GetTotalClose(ticketId string) (int64, error) {
// 	var ticket model.Ticket
// 	var count int64
// 	// nanti diganti ticket_prioritynya
// 	err := t.db.Find(&ticket).Where("ticket_id = ? and ticket_priority = 'close'", ticketId).Count(&count).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return 0, err
// 		}
// 		fmt.Println(err.Error())
// 		return 0, fmt.Errorf("komplain dengan id %v tidak ditemukan", ticketId)
// 	}
// 	return count, nil
// }

// func (t *ticketRepository) GetTotalProgress(ticketId string) (int64, error) {
// 	var ticket model.Ticket
// 	var count int64
// 	// nanti diganti ticket_prioritynya
// 	err := t.db.Find(&ticket).Where("ticket_id = ? and ticket_priority = 'progress'", ticketId).Count(&count).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return 0, err
// 		}
// 		fmt.Println(err.Error())
// 		return 0, fmt.Errorf("komplain dengan id %v tidak ditemukan", ticketId)
// 	}
// 	return count, nil
// }

// func (t *ticketRepository) UpdatePIC(ticket *model.JsonTicket, userId string) error {
// 	err := t.db.Model(model.Ticket{}).Where("ticket_id = ?", ticket.ID).Update("pic", model.Pic{UserId: userId}).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func NewTicketRepository(db *gorm.DB) TicketRepository {
	repo := new(ticketRepository)
	repo.db = db
	return repo
}
