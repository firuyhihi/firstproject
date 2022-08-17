package model

import "encoding/json"

type UserRole struct {
	Id     int    `gorm:"primaryKey" json:"userRoleId"`
	UserId string `json:"userId"`
	RoleId int    `json:"roleId"`
	Role   Role   `gorm:"foreignKey:RoleId; references:Id" json:"role"`
}

func (UserRole) TableName() string {
	return "u_user_role"
}

func (u *UserRole) ToJSON() string {
	userRole, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		return ""
	}
	return string(userRole)
}
