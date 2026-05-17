package dto

import (
	"hta-platform/internal/media/domain/model/entity"
)

type GetChapterImagesReq struct {
	MediaURL string `form:"media_url" binding:"required" validate:"required" dc:"Media URL"`
}

type Image struct {
	Url         string `json:"url" binding:"required" validate:"required" dc:"URL ảnh"`
	Description string `json:"description" binding:"omitempty" validate:"omitempty" dc:"Mô tả ảnh"`
	Source      string `json:"source" binding:"omitempty" validate:"omitempty" dc:"Nguồn ảnh"`
}

type ChapterImageReq struct {
	Order  int64   `json:"order" binding:"required" validate:"required" dc:"Thứ tự chapter"`
	Images []Image `json:"images" binding:"required" validate:"required" dc:"Danh sách ảnh"`
}

type CreateChapterImageReq struct {
	MediaUrl      string            `json:"media_url" binding:"required" validate:"required" dc:"Media URL"`
	ChapterOrder  int64             `json:"chapter_order" binding:"required" validate:"required" dc:"Thứ tự chapter"`
	ChapterImages []ChapterImageReq `json:"chapter_images" binding:"required" validate:"required" dc:"Danh sách ảnh"`
}

type CreateMediaChapterReq struct {
	MediaUrl string             `json:"media_url" binding:"required" validate:"required" dc:"Media URL"`
	Chapters []CreateChapterReq `json:"chapters" binding:"required" validate:"required" dc:"Danh sách chapter"`
}

type CreateChapterReq struct {
	Name   string `json:"name" binding:"required" validate:"required,min=1" dc:"Tên chapter"`
	Order  int64  `json:"order" binding:"required" validate:"required" dc:"Thứ tự chapter"`
	Source string `json:"source" binding:"required" validate:"required" dc:"Nguồn chapter"`
}

type MediaChapterRes struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	URL      string            `json:"url"`
	Language string            `json:"language"`
	Order    int64             `json:"order"`
	Source   string            `json:"source"`
	Images   []ChapterImageRes `json:"images,omitempty"`
}

func (r *MediaChapterRes) SetData(chapter entity.MediaChapter) {
	r.ID = chapter.ID.String()
	r.Name = chapter.Name
	r.URL = chapter.URL
	r.Language = chapter.Language
	r.Order = chapter.Order
	r.Source = chapter.Source
	if len(chapter.Images) > 0 {
		r.Images = make([]ChapterImageRes, len(chapter.Images))
		for i, img := range chapter.Images {
			r.Images[i].SetData(img)
		}
	}
}

type ChapterImageRes struct {
	ID     string  `json:"id"`
	Order  int64   `json:"order"`
	Images []Image `json:"images"`
}

func (r *ChapterImageRes) SetData(image entity.ChapterImage) {
	r.ID = image.ID.String()
	r.Order = image.Order
	if len(image.Images) > 0 {
		r.Images = make([]Image, len(image.Images))
		for i, img := range image.Images {
			r.Images[i] = Image{
				Url:         img.URL,
				Description: img.Description,
				Source:      img.Source,
			}
		}
	}
}
