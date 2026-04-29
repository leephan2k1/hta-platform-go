package entity

import (
	imageEntity "hta-platform/internal/image/domain/model/entity"
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type ChapterImage struct {
	base.CommonModel
	Order     int64               `gorm:"column:order;uniqueIndex:idx_order_chapter" json:"order"`
	ChapterID uuid.UUID           `gorm:"column:chapter_id;type:uuid;not null;uniqueIndex:idx_order_chapter" json:"chapter_id"`
	Images    []imageEntity.Image `gorm:"foreignKey:ResourceID;references:ID" json:"images,omitempty"`
}

func (ChapterImage) TableName() string {
	return "hta.chapter_image"
}
