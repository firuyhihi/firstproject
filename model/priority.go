package model

import "encoding/json"

type Priority struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	PriorityName string `gorm:"not null type:varchar(50)" json:"priorityName"`
}

func (Priority) TableName() string {
	return "t_priority"
}

func (p *Priority) ToString() string {
	priority, err := json.MarshalIndent(p, "", "")
	if err != nil {
		return ""
	}
	return string(priority)
}
