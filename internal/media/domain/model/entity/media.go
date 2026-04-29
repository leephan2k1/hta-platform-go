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
	Status      Status                    `gorm:"foreignKey:StatusID;references:ID" json:"status,omitempty"`
	TypeID      uuid.UUID                 `gorm:"column:type_id;type:uuid" json:"type_id"`
	Type        Type                      `gorm:"foreignKey:TypeID;references:ID" json:"type,omitempty"`
	IsNSFW      bool                      `gorm:"column:is_nsfw;default:false" json:"is_nsfw"`
	Thumbnail   string                    `gorm:"column:thumbnail" json:"thumbnail"`
	Source      string                    `gorm:"column:source" json:"source"`
	SysStatus   string                    `gorm:"column:sys_status;type:varchar(32);default:active" json:"sys_status"`
	Categories  []categoryEntity.Category `gorm:"many2many:hta.media_to_category;foreignKey:ID;joinForeignKey:MediaID;References:ID;joinReferences:CategoryID" json:"categories,omitempty"`
	Authors     []authorEntity.Author     `gorm:"many2many:hta.media_to_author;foreignKey:ID;joinForeignKey:MediaID;References:ID;joinReferences:AuthorID" json:"authors,omitempty"`
	OtherNames  []MediaOtherName          `gorm:"foreignKey:MediaID;references:ID" json:"other_names,omitempty"`
	Chapters    []MediaChapter            `gorm:"foreignKey:MediaID;references:ID" json:"chapters,omitempty"`
}

func (Media) TableName() string {
	return "hta.media"
}
