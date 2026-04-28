package entity

import (
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type ChapterImage struct {
	base.CommonModel
	Order     int64     `gorm:"column:order;uniqueIndex:idx_order_chapter" json:"order"`
	ChapterID uuid.UUID `gorm:"column:chapter_id;type:uuid;not null;uniqueIndex:idx_order_chapter" json:"chapter_id"`
}

func (ChapterImage) TableName() string {
	return "hta.chapter_image"
}
