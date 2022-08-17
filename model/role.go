package model

import "encoding/json"

type Role struct {
	Id       int    `gorm:"primaryKey" json:"roleId"`
	RoleName string `gorm:"column:role_name" json:"roleName"`
	RoleCode string `gorm:"column:role_code" json:"roleCode"`
}

func (Role) TableName() string {
	return "u_role"
}

func (r *Role) ToJSON() string {
	role, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return ""
	}
	return string(role)
}
