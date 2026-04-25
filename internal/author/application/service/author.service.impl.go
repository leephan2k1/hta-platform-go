package service

import (
	"context"
	"hta-platform/internal/author/controller/dto"
	"hta-platform/internal/author/domain/model/entity"
	"hta-platform/internal/author/domain/repository"

	"github.com/gosimple/slug"
)

type authorService struct {
	authorRepo repository.AuthorRepository
}

// CreateAuthor implements [AuthorService].
func (a *authorService) CreateAuthor(ctx context.Context, req *dto.AuthorReq) (entity.Author, error) {
	author := entity.Author{
		Name:      req.Name,
		AuthorURL: slug.Make(req.Name),
	}
	return a.authorRepo.CreateAuthor(ctx, &author)
}

// GetAuthorByUrl implements [AuthorService].
func (a *authorService) GetAuthorByUrl(ctx context.Context, url string) (*entity.Author, error) {
	return a.authorRepo.FindAuthorByUrl(ctx, url)
}

func NewAuthorService(authorRepo repository.AuthorRepository) AuthorService {
	return &authorService{authorRepo: authorRepo}
}
