package entity

import (
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type Image struct {
	base.BaseModel
	URL         string    `gorm:"column:url;unique;not null" json:"url"`
	Description string    `gorm:"column:description" json:"description"`
	ResourceID  uuid.UUID `gorm:"column:resource_id;type:uuid;not null" json:"resource_id"`
	Source      string    `gorm:"column:source" json:"source"`
}

func (Image) TableName() string {
	return "hta.image"
}
