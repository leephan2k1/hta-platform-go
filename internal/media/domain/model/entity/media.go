package entity

import (
	authorEntity "hta-platform/internal/author/domain/model/entity"
	categoryEntity "hta-platform/internal/category/domain/model/entity"
	"hta-platform/pkg/base"

	"github.com/google/uuid"
)

type Media struct {
	base.CommonModel
	Name        string                    `gorm:"column:name;unique;not null" json:"name"`
	Description string                    `gorm:"column:description" json:"description"`
	URL         string                    `gorm:"column:url;unique;not null" json:"url"`
	StatusID    uuid.UUID                 `gorm:"column:status_id;type:uuid" json:"status_id"`
	TypeID      uuid.UUID                 `gorm:"column:type_id;type:uuid" json:"type_id"`
	IsNSFW      bool                      `gorm:"column:is_nsfw;default:false" json:"is_nsfw"`
	Thumbnail   string                    `gorm:"column:thumbnail" json:"thumbnail"`
	Source      string                    `gorm:"column:source" json:"source"`
	Categories  []categoryEntity.Category `gorm:"many2many:hta.media_to_category;foreignKey:ID;joinForeignKey:MediaID;References:ID;joinReferences:CategoryID" json:"categories,omitempty"`
	Authors     []authorEntity.Author     `gorm:"many2many:hta.media_to_author;foreignKey:ID;joinForeignKey:MediaID;References:ID;joinReferences:AuthorID" json:"authors,omitempty"`
}

func (Media) TableName() string {
	return "hta.media"
}
