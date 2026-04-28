package dto

import (
	"hta-platform/internal/media/domain/model/entity"
)

type CreateMediaChapterReq struct {
	MediaUrl string             `json:"media_url" binding:"required" validate:"required" dc:"Media URL"`
	Chapters []CreateChapterReq `json:"chapters" binding:"required" validate:"required" dc:"Danh sách chapter"`
}

type CreateChapterReq struct {
	Name  string `json:"name" binding:"required" validate:"required,min=1" dc:"Tên chapter"`
	Order int64  `json:"order" binding:"required" validate:"required" dc:"Thứ tự chapter"`
}

type MediaChapterRes struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	URL      string            `json:"url"`
	Language string            `json:"language"`
	Order    int64             `json:"order"`
	Images   []ChapterImageRes `json:"images,omitempty"`
}

func (r *MediaChapterRes) SetData(chapter entity.MediaChapter) {
	r.ID = chapter.ID.String()
	r.Name = chapter.Name
	r.URL = chapter.URL
	r.Language = chapter.Language
	r.Order = chapter.Order
	if len(chapter.Images) > 0 {
		r.Images = make([]ChapterImageRes, len(chapter.Images))
		for i, img := range chapter.Images {
			r.Images[i].SetData(img)
		}
	}
}

type ChapterImageRes struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Order int64  `json:"order"`
}

func (r *ChapterImageRes) SetData(image entity.ChapterImage) {
	r.ID = image.ID.String()
	r.URL = image.URL
	r.Order = image.Order
}
