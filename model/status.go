package model

import "encoding/json"

type Status struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	StatusName string `gorm:"not null type:varchar(50)" json:"statusName"`
}

func (Status) TableName() string {
	return "t_status"
}

func (s *Status) ToString() string {
	status, err := json.MarshalIndent(s, "", "")
	if err != nil {
		return ""
	}
	return string(status)
}
