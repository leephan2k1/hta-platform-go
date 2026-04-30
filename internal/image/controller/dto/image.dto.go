package dto

type MigrateThumbnailReq struct {
	Source string `form:"source" binding:"required" validate:"required" dc:"Nguồn"`

	Description string `form:"description" binding:"required" validate:"required" dc:"Mô tả"`
}

type StreamImageReq struct {
	Source string `form:"source" binding:"required" validate:"required" dc:"Nguồn"`

	URL string `form:"url" binding:"required" validate:"required" dc:"URL"`
}
