package repository

import (
	"context"
	"hta-platform/internal/author/domain/model/entity"
	"hta-platform/internal/author/domain/repository"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type authorRepository struct {
	DB *gorm.DB
}

// CreateAuthor implements [repository.AuthorRepository].
func (a *authorRepository) CreateAuthor(ctx context.Context, author *entity.Author) (entity.Author, error) {
	result := a.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "author_url"}},
		DoUpdates: clause.AssignmentColumns([]string{"author_url", "name", "updated_at"}),
	}).Create(author)

	if result.Error != nil {
		return entity.Author{}, result.Error
	}
	return *author, nil
}

// FindAuthorByUrl implements [repository.AuthorRepository].
func (a *authorRepository) FindAuthorByUrl(ctx context.Context, authorURL string) (*entity.Author, error) {
	var author entity.Author
	result := a.DB.Where("author_url = ?", authorURL).First(&author)
	if result.Error != nil {
		return nil, result.Error
	}
	return &author, nil
}

func (a *authorRepository) FindAuthors(ctx context.Context, name string, limit, offset int) ([]entity.Author, int64, error) {
	var authors []entity.Author
	var total int64
	query := a.DB.Model(&entity.Author{})
	if name != "" {
		query = query.Where("LOWER(name) ILIKE ?", "%"+strings.ToLower(name)+"%")
	}

	query = query.Order("name ASC")

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Limit(limit).Offset(offset).Find(&authors).Error
	return authors, total, err
}

func NewAuthorRepository(db *gorm.DB) repository.AuthorRepository {
	return &authorRepository{db}
}
