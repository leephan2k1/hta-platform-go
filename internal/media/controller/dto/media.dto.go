package dto

import (
	authorDto "hta-platform/internal/author/controller/dto"
	categoryDto "hta-platform/internal/category/controller/dto"
	"hta-platform/pkg/base"
	"strings"
	"time"
)

type MediaResponse struct {
	ID          string                    `json:"id"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	URL         string                    `json:"url"`
	StatusID    string                    `json:"status_id"`
	TypeID      string                    `json:"type_id"`
	IsNSFW      bool                      `json:"is_nsfw"`
	Thumbnail   string                    `json:"thumbnail"`
	Source      string                    `json:"source"`
	Categories  []categoryDto.CategoryRes `json:"categories,omitempty"`
	Authors     []authorDto.AuthorRes     `json:"authors,omitempty"`
}

type GetMediasReq struct {
	Name       string   `form:"name"`
	SortBy     []string `form:"sortBy"`
	Authors    []string `form:"authors"`
	Categories []string `form:"categories"`
	IsNSFW     bool     `form:"isNSFW"`
	Limit      int      `form:"limit"`
	Page       int      `form:"page"`
}

type GetMediasRes struct {
	base.BasePaginationRes
	Items []MediaResponse `json:"items"`
}

func (r *GetMediasReq) Normalize() {
	r.Name = strings.TrimSpace(r.Name)
}

type CreateMediaReq struct {
	Name              string   `json:"name" binding:"required" validate:"required,min=1" dc:"Tên media"`
	Thumbnail         string   `json:"thumbnail" validate:"omitempty,url" dc:"Thumbnail"`
	Description       string   `json:"description" validate:"omitempty" dc:"Mô tả"`
	StatusID          string   `json:"statusId" binding:"required" validate:"required,uuid" dc:"ID trạng thái"`
	TypeID            string   `json:"typeId" binding:"required" validate:"required,uuid" dc:"ID loại"`
	IsNSFW            bool     `json:"isNSFW" dc:"Gán nhãn NSFW"`
	AuthorIDs         []string `json:"authorIds" validate:"omitempty,dive,uuid" dc:"Danh sách ID tác giả"`
	CategoryIDs       []string `json:"categoryIds" validate:"omitempty,dive,uuid" dc:"Danh sách ID danh mục"`
	OtherName         string   `json:"otherName" validate:"omitempty" dc:"Tên khác"`
	OtherNameLanguage string   `json:"otherNameLanguage" validate:"omitempty" dc:"Ngôn ngữ tên khác"`
}

func (r *CreateMediaReq) Normalize() {
	r.Name = strings.TrimSpace(r.Name)
}
