package entity

import (
	"hta-platform/pkg/base"
)

type Author struct {
	base.CommonModel
	Name      string `gorm:"column:name;unique;not null" json:"name"`
	AuthorURL string `gorm:"column:author_url;unique" json:"author_url"`
}

func (Author) TableName() string {
	return "hta.author"
}
