package domain

import (
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type MediaChapter struct {
	base.CommonModel
	Name     string    `gorm:"column:name;unique;not null" json:"name"`
	URL      string    `gorm:"column:url;unique" json:"url"`
	Language string    `gorm:"column:language" json:"language"`
	Order    int64     `gorm:"column:order" json:"order"`
	MediaID  uuid.UUID `gorm:"column:media_id;type:uuid;not null" json:"media_id"`
	Source   string    `gorm:"column:source" json:"source"`
}

func (MediaChapter) TableName() string {
	return "hta.media_chapter"
}
