package dto

import (
	entity "hta-platform/internal/author/domain/model/entity"
	"strings"
)

type AuthorReq struct {
	Name string `json:"name" binding:"required" validate:"required,min=1" dc:"Tên tác giả"`
}

type AuthorRes struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AuthorURL string `json:"author_url"`
}

func (r *AuthorRes) SetData(author entity.Author) {
	r.ID = author.ID.String()
	r.Name = author.Name
	r.AuthorURL = author.AuthorURL
}

func (r *AuthorReq) Normalize() {
	r.Name = strings.TrimSpace(r.Name)
}
