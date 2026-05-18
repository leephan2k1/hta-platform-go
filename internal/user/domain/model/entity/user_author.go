package entity

import (
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type UserAuthor struct {
	base.BaseModel

	UserID   string    `gorm:"column:user_id;type:text" json:"user_id"`
	AuthorID uuid.UUID `gorm:"column:author_id;type:uuid" json:"author_id"`
}

func (m *UserAuthor) TableName() string {
	return "hta.user_to_author"
}
