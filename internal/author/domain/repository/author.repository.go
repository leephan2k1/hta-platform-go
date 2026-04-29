package repository

import (
	"context"

	authorEntity "hta-platform/internal/author/domain/model/entity"
)

type AuthorRepository interface {
	CreateAuthor(ctx context.Context, author *authorEntity.Author) (authorEntity.Author, error)

	FindAuthorByUrl(ctx context.Context, authorURL string) (*authorEntity.Author, error)

	FindAuthors(ctx context.Context, name string, limit, offset int) ([]authorEntity.Author, int64, error)
}
