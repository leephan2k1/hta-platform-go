package dto

type StreamImageReq struct {
	Source string `form:"source" binding:"required" validate:"required" dc:"Nguồn"`

	URL string `form:"url" binding:"required" validate:"required" dc:"URL"`
}
