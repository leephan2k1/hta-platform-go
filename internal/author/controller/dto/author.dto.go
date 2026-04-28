package dto

import "strings"

type AuthorReq struct {
	Name string `json:"name" binding:"required" validate:"required,min=1" dc:"Tên tác giả"`
}

type AuthorRes struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AuthorURL string `json:"author_url"`
}

func (r *AuthorReq) Normalize() {
	r.Name = strings.TrimSpace(r.Name)
}
