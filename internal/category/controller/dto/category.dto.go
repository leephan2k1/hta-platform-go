package dto

import (
	entity "hta-platform/internal/category/domain/model/entity"
	"strings"
)

type CategoryReq struct {
	Name string `json:"name" binding:"required" validate:"required,min=1" dc:"Tên danh mục"`
}

type CategoryRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (r *CategoryRes) SetData(category entity.Category) {
	r.ID = category.ID.String()
	r.Name = category.Name
	r.Slug = category.Slug
}

func (r *CategoryReq) Normalize() {
	r.Name = strings.TrimSpace(r.Name)
}
