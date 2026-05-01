package entity

import (
	"hta-platform/pkg/base"
	"time"

	"github.com/google/uuid"
)

type UserReadingSession struct {
	base.BaseModel

	UserID    string    `gorm:"column:user_id;type:text;index" json:"user_id"`
	ChapterID uuid.UUID `gorm:"column:chapter_id;type:uuid;index" json:"chapter_id"`
	MediaID   uuid.UUID `gorm:"column:media_id;type:uuid;index" json:"media_id"`

	StartedAt time.Time `gorm:"column:started_at;type:timestamp with time zone" json:"started_at"`
	EndedAt   time.Time `gorm:"column:ended_at;type:timestamp with time zone" json:"ended_at"`
	Duration  int64     `gorm:"column:duration;type:bigint" json:"duration"`
}

func (m *UserReadingSession) TableName() string {
	return "hta.user_reading_sessions"
}
