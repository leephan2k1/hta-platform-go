package entity

import (
	"hta-platform/pkg/base"
)

type Type struct {
	base.BaseModel
	Name string `gorm:"column:name" json:"name"`
}

func (Type) TableName() string {
	return "hta.type"
}
