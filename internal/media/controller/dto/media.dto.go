package dto

import "strings"

type MediaReq struct {
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

func (r *MediaReq) Normalize() {
	r.Name = strings.TrimSpace(r.Name)
}
