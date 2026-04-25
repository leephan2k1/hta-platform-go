package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaToAuthor struct {
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:LOCALTIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:LOCALTIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
	CreatedBy string         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy string         `gorm:"column:deleted_by" json:"deleted_by"`
	MediaID   uuid.UUID      `gorm:"column:media_id;primaryKey;type:uuid" json:"media_id"`
	AuthorID  uuid.UUID      `gorm:"column:author_id;primaryKey;type:uuid" json:"author_id"`
}

func (MediaToAuthor) TableName() string {
	return "hta.media_to_author"
}
