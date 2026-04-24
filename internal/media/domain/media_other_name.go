package domain

import (
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type MediaOtherName struct {
	base.CommonModel
	Name     string    `gorm:"column:name;unique;not null" json:"name"`
	Language string    `gorm:"column:language" json:"language"`
	MediaID  uuid.UUID `gorm:"column:media_id;type:uuid;not null" json:"media_id"`
}

func (MediaOtherName) TableName() string {
	return "hta.media_other_name"
}
