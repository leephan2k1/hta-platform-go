package service

import (
	"context"
	"hta-platform/internal/author/controller/dto"
	"hta-platform/internal/author/domain/model/entity"
)

type AuthorService interface {
	CreateAuthor(ctx context.Context, req *dto.AuthorReq) (entity.Author, error)

	GetAuthorByUrl(ctx context.Context, url string) (*entity.Author, error)

	GetAuthors(ctx context.Context, req *dto.GetAuthorsReq) (*dto.GetAuthorsRes, error)
}
