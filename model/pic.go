package model

import "encoding/json"

type Pic struct {
	Id            int           `gorm:"primaryKey" json:"picId"`
	RoleId        int           `json:"picRole"`
	UserId        string        `json:"picUser"`
	DepartmentId  int           `json:"picDepartment"`
	PicDepartment PicDepartment `gorm:"foreignKey:DepartmentId; references:Id"`
}

func (Pic) TableName() string {
	return "u_pic"
}

func (p *Pic) ToJSON() string {
	pic, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return ""
	}
	return string(pic)
}
