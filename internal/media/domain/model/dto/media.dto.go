package dto

import (
	"hta-platform/internal/media/domain/model/entity"
)

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
	Total int64          `json:"total"`
	Items []entity.Media `json:"items"`
}
