package dto

import "strings"

type UserReadingSessionRes struct {
	MediaID     string `json:"mediaId"`
	Duration    int64  `json:"duration"`
	FirstReadAt string `json:"firstReadAt"`
	LastReadAt  string `json:"lastReadAt"`
}

type UserReadingSessionStartRes struct {
	SessionID string `json:"sessionId"`
}

type UserReadingSessionReq struct {
	MediaID string `json:"mediaId" validate:"required,min=1"`
}

type UserReadingSessionEndReq struct {
	SessionID string `json:"sessionId" validate:"required,min=1"`
}

type UserReadingProgressReq struct {
	ChapterID  string `json:"chapterId" validate:"required,min=1"`
	MediaID    string `json:"mediaId" validate:"required,min=1"`
	ImageOrder *int   `json:"imageOrder" validate:"required"`
}

type UserToResourceReq struct {
	ResourceID string `json:"resourceId" validate:"required,min=1"`
}

type RegisterUserReq struct {
	Auth0Id    string `json:"auth0Id" validate:"required,min=1"`
	Email      string `json:"email" validate:"required,min=1,email"`
	Picture    string `json:"picture" validate:"omitempty"`
	GivenName  string `json:"givenName" validate:"omitempty"`
	FamilyName string `json:"familyName" validate:"omitempty"`
}

func (r *RegisterUserReq) Normalize() {
	r.Auth0Id = strings.TrimSpace(r.Auth0Id)
	r.Email = strings.TrimSpace(r.Email)
	r.Picture = strings.TrimSpace(r.Picture)
	r.GivenName = strings.TrimSpace(r.GivenName)
	r.FamilyName = strings.TrimSpace(r.FamilyName)
}
