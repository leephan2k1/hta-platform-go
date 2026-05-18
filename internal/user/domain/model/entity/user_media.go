package entity

import (
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type UserMedia struct {
	base.BaseModel

	UserID  string    `gorm:"column:user_id;type:text" json:"user_id"`
	MediaID uuid.UUID `gorm:"column:media_id;type:uuid" json:"media_id"`
}

func (m *UserMedia) TableName() string {
	return "hta.user_to_media"
}
