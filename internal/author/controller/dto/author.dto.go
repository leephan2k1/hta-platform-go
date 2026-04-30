package dto

import (
	entity "hta-platform/internal/author/domain/model/entity"
	imageEntity "hta-platform/internal/image/domain/model/entity"
	mediaEntity "hta-platform/internal/media/domain/model/entity"
	"hta-platform/pkg/base"
	"strings"
)

type GetAuthorsRes struct {
	base.BasePaginationRes
	Items []AuthorRes `json:"items"`
}

type GetAuthorsReq struct {
	Name  string `form:"name" dc:"Tên tác giả"`
	Page  int    `form:"page" dc:"Trang"`
	Limit int    `form:"limit" dc:"Giới hạn"`
}

type AuthorReq struct {
	Name string `json:"name" binding:"required" validate:"required,min=1" dc:"Tên tác giả"`
}

type AuthorRes struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	AuthorURL  string          `json:"author_url"`
	FirstMedia *AuthorMediaRes `json:"first_media,omitempty"`
}

type AuthorMediaRes struct {
	Name   string                `json:"name"`
	Images []AuthorMediaImageRes `json:"images,omitempty"`
}

type AuthorMediaImageRes struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Source      string `json:"source"`
}

func (r *AuthorRes) SetData(author entity.Author) {
	r.ID = author.ID.String()
	r.Name = author.Name
	r.AuthorURL = author.AuthorURL
}

func (r *AuthorMediaRes) SetData(media mediaEntity.Media) {
	r.Name = media.Name
	if len(media.Images) > 0 {
		r.Images = make([]AuthorMediaImageRes, len(media.Images))
		for i, img := range media.Images {
			r.Images[i].SetData(img)
		}
	}
}

func (r *AuthorMediaImageRes) SetData(img imageEntity.Image) {
	r.ID = img.ID.String()
	r.URL = img.URL
	r.Description = img.Description
	r.Source = img.Source
}

func (r *AuthorReq) Normalize() {
	r.Name = strings.TrimSpace(r.Name)
}
