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

func (a *authorService) GetAuthors(ctx context.Context, req *dto.GetAuthorsReq) (*dto.GetAuthorsRes, error) {
	limit := req.Limit
	if limit <= 0 {
		limit = 20
	}
	page := req.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	authors, total, err := a.authorRepo.FindAuthors(ctx, req.Name, limit, offset)
	if err != nil {
		return nil, err
	}

	var res dto.GetAuthorsRes
	res.SetPagination(total, req.Page, req.Limit)

	res.Items = make([]dto.AuthorRes, len(authors))
	for i, author := range authors {
		res.Items[i].SetData(author)
	}

	return &res, nil
}

func NewAuthorService(authorRepo repository.AuthorRepository) AuthorService {
	return &authorService{authorRepo: authorRepo}
}
