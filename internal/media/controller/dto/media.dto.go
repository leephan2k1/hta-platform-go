package dto

import (
	authorDto "hta-platform/internal/author/controller/dto"
	categoryDto "hta-platform/internal/category/controller/dto"
	imageEntity "hta-platform/internal/image/domain/model/entity"
	"hta-platform/internal/media/domain/model/entity"
	"hta-platform/pkg/base"
	"strings"
	"time"

	"github.com/google/uuid"
)

type MediaResponse struct {
	ID          string                    `json:"id"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	URL         string                    `json:"url"`
	StatusID    string                    `json:"status_id"`
	Status      *StatusRes                `json:"status,omitempty"`
	TypeID      string                    `json:"type_id"`
	Type        *TypeRes                  `json:"type,omitempty"`
	IsNSFW      bool                      `json:"is_nsfw"`
	Thumbnail   string                    `json:"thumbnail"`
	Source      string                    `json:"source"`
	SysStatus   string                    `json:"sys_status"`
	Categories  []categoryDto.CategoryRes `json:"categories,omitempty"`
	Authors     []authorDto.AuthorRes     `json:"authors,omitempty"`
	OtherNames  []OtherNameRes            `json:"other_names,omitempty"`
	Chapters    []ChapterRes              `json:"chapters,omitempty"`
	Images      []ImageRes                `json:"images,omitempty"`
	Progress    float64                   `json:"progress"`
}

func (r *MediaResponse) SetData(media entity.Media) {
	r.ID = media.ID.String()
	r.CreatedAt = media.CreatedAt
	r.UpdatedAt = media.UpdatedAt
	r.Name = media.Name
	r.Description = media.Description
	r.URL = media.URL
	r.StatusID = media.StatusID.String()
	r.TypeID = media.TypeID.String()
	r.IsNSFW = media.IsNSFW
	r.Thumbnail = media.Thumbnail
	r.Source = media.Source
	r.SysStatus = media.SysStatus

	if media.Status.ID != uuid.Nil {
		r.Status = &StatusRes{}
		r.Status.SetData(media.Status)
	}

	if media.Type.ID != uuid.Nil {
		r.Type = &TypeRes{}
		r.Type.SetData(media.Type)
	}

	if len(media.Categories) > 0 {
		r.Categories = make([]categoryDto.CategoryRes, len(media.Categories))
		for i, cat := range media.Categories {
			r.Categories[i].SetData(cat)
		}
	}

	if len(media.Authors) > 0 {
		r.Authors = make([]authorDto.AuthorRes, len(media.Authors))
		for i, author := range media.Authors {
			r.Authors[i].SetData(author)
		}
	}

	if len(media.OtherNames) > 0 {
		r.OtherNames = make([]OtherNameRes, len(media.OtherNames))
		for i, on := range media.OtherNames {
			r.OtherNames[i].SetData(on)
		}
	}

	if len(media.Chapters) > 0 {
		r.Chapters = make([]ChapterRes, len(media.Chapters))
		for i, ch := range media.Chapters {
			r.Chapters[i].SetData(ch)
		}
	}

	if len(media.Images) > 0 {
		r.Images = make([]ImageRes, len(media.Images))
		for i, img := range media.Images {
			r.Images[i].SetData(img)
		}
	}
}

type StatusRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (r *StatusRes) SetData(status entity.Status) {
	r.ID = status.ID.String()
	r.Name = status.Name
}

type TypeRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (r *TypeRes) SetData(mediaType entity.Type) {
	r.ID = mediaType.ID.String()
	r.Name = mediaType.Name
}

type OtherNameRes struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

func (r *OtherNameRes) SetData(on entity.MediaOtherName) {
	r.Name = on.Name
	r.Language = on.Language
}

type ChapterRes struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Language string `json:"language"`
	Order    int64  `json:"order"`
}

func (r *ChapterRes) SetData(chapter entity.MediaChapter) {
	r.ID = chapter.ID.String()
	r.Name = chapter.Name
	r.URL = chapter.URL
	r.Language = chapter.Language
	r.Order = chapter.Order
}

type ImageRes struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Source      string `json:"source"`
}

func (r *ImageRes) SetData(img imageEntity.Image) {
	r.ID = img.ID.String()
	r.URL = img.URL
	r.Description = img.Description
	r.Source = img.Source
}

type GetMediasReq struct {
	Name       string   `form:"name"`
	SortBy     []string `form:"sortBy"`
	Authors    []string `form:"authors"`
	Categories []string `form:"categories"`
	IsNSFW     bool     `form:"isNSFW"`
	SysStatus  string   `form:"sysStatus"`
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
