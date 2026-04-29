package dto

import (
	entity "hta-platform/internal/author/domain/model/entity"
	"hta-platform/pkg/base"
	"strings"
)

type GetAuthorsRes struct {
	base.BasePaginationRes
	Items []AuthorRes `json:"items"`
}

type GetAuthorsReq struct {
	Name  string `query:"name" dc:"Tên tác giả"`
	Page  int    `query:"page" dc:"Trang"`
	Limit int    `query:"limit" dc:"Giới hạn"`
}

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
