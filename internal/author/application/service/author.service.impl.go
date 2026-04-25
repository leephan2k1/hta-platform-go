package service

import (
	"context"
	"hta-platform/internal/author/controller/dto"
	"hta-platform/internal/author/domain/model/entity"
	"hta-platform/internal/author/domain/repository"
)

type authorService struct {
	authorRepo repository.AuthorRepository
}

// CreateAuthor implements [AuthorService].
func (a *authorService) CreateAuthor(ctx context.Context, req *dto.AuthorReq) error {
	panic("unimplemented")
}

// GetAuthorByUrl implements [AuthorService].
func (a *authorService) GetAuthorByUrl(ctx context.Context, url string) (*entity.Author, error) {
	panic("unimplemented")
}

func NewAuthorService(authorRepo repository.AuthorRepository) AuthorService {
	return &authorService{authorRepo: authorRepo}
}
