package repository

import (
	"context"

	authorEntity "hta-platform/internal/author/domain/model/entity"
)

type AuthorRepository interface {
	CreateAuthor(ctx context.Context, author *authorEntity.Author) (authorEntity.Author, error)

	FindAuthorByUrl(ctx context.Context, authorURL string) (*authorEntity.Author, error)
}
