package entity

import (
	"hta-platform/pkg/base"
)

type Category struct {
	base.CommonModel
	Name string `gorm:"column:name;unique;not null" json:"name"`
	Slug string `gorm:"column:slug;unique;not null" json:"slug"`
}

func (Category) TableName() string {
	return "hta.category"
}
