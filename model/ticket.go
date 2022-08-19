package model

import (
	"encoding/json"
	"time"


)

type Ticket struct {
	TicketId         string        `gorm:"primaryKey;type:varchar(8)" json:"ticketId"`
	CsId             string        `gorm:"not null; type:varchar(5)" json:"csId"`
	TicketSubject    string        `gorm:"not null; type:varchar(500)" json:"ticketSubject"`
	DepartmentId     int           `gorm:"not null" json:"departmentId"`
	TicketMessage    string        `gorm:"not null; type:varchar(1000)" json:"ticketMessage"`
	PriorityId       int           `gorm:"not null" json:"priorityId"`
	TicketAttachment string        `gorm:"type:varchar(1000)" json:"ticketAttachment"`
	PicId            int           `gorm:"not null" json:"picId"`
	StatusId         int           `gorm:"not null" json:"statusId"`
	CreatedAt        time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"ticketDate"`
	UpdatedAt        time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	IsActive         bool          `gorm:"type:boolean; column:is_active default: true" json:"isActive"`
	Priority         Priority      `gorm:"foreignKey:PriorityId; references:Id"`
	Status           Status        `gorm:"foreignKey:StatusId; references:Id"`
	Pic              Pic           `gorm:"foreignKey:PicId; references:Id"`
	PicDepartment    PicDepartment `gorm:"foreignKey:DepartmentId; refernces:Id"`
}

type JsonTicket struct {
	ID string `json:"id"`
}

func (Ticket) TableName() string {
	return "t_ticket"
}

func (t *Ticket) ToString() string {
	ticket, err := json.MarshalIndent(t, "", "")
	if err != nil {
		return ""
	}
	return string(ticket)
}
