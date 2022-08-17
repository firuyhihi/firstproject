package model

import "encoding/json"

type PicDepartment struct {
	Id             int    `gorm:"primaryKey" json:"picDepartmentId"`
	DepartmentName string `gorm:"column:department_name" json:"picDepartmentName"`
	DepartmentCode string `gorm:"column:department_code" json:"picDepartmentCode"`
}

func (PicDepartment) TableName() string {
	return "u_pic_department"
}

func (p *PicDepartment) ToJSON() string {
	picDepartment, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return ""
	}
	return string(picDepartment)
}
