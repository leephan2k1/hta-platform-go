package domain

import (
	"hta-platform/pkg/base"
)

type Status struct {
	base.BaseModel
	Name string `gorm:"column:name" json:"name"`
}

func (Status) TableName() string {
	return "hta.status"
}
