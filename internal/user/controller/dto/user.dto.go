package dto

import "strings"

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
