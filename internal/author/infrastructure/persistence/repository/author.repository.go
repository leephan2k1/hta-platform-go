package repository

import (
	"context"
	"hta-platform/internal/author/domain/model/entity"
	"hta-platform/internal/author/domain/repository"

	"gorm.io/gorm"
)

type authorRepository struct {
	DB *gorm.DB
}

// CreateAuthor implements [repository.AuthorRepository].
func (a *authorRepository) CreateAuthor(ctx context.Context, author *entity.Author) error {
	panic("unimplemented")
}

// FindAuthorByUrl implements [repository.AuthorRepository].
func (a *authorRepository) FindAuthorByUrl(ctx context.Context, authorURL string) (*entity.Author, error) {
	panic("unimplemented")
}

func NewAuthorRepository(db *gorm.DB) repository.AuthorRepository {
	return &authorRepository{db}
}
