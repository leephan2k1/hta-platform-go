package entity

import (
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type MediaChapter struct {
	base.CommonModel
	Name     string         `gorm:"column:name;not null" json:"name"`
	URL      string         `gorm:"column:url;uniqueIndex:idx_media_chapter_unique" json:"url"`
	Language string         `gorm:"column:language" json:"language"`
	Order    int64          `gorm:"column:order;uniqueIndex:idx_media_chapter_unique" json:"order"`
	MediaID  uuid.UUID      `gorm:"column:media_id;type:uuid;not null;uniqueIndex:idx_media_chapter_unique" json:"media_id"`
	Source   string         `gorm:"column:source" json:"source"`
	Images   []ChapterImage `gorm:"foreignKey:ChapterID;references:ID" json:"images,omitempty"`
}

func (MediaChapter) TableName() string {
	return "hta.media_chapter"
}
