package entity

import (
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type UserReadingProgress struct {
	base.BaseModel

	UserID            string    `gorm:"column:user_id;type:text;index" json:"user_id"`
	ChapterID         uuid.UUID `gorm:"column:chapter_id;type:uuid;index" json:"chapter_id"`
	ChapterImageOrder int       `gorm:"column:chapter_image_order;type:integer" json:"chapter_image_order"`
	MediaID           uuid.UUID `gorm:"column:media_id;type:uuid;index" json:"media_id"`
}

func (m *UserReadingProgress) TableName() string {
	return "hta.user_reading_progress"
}
