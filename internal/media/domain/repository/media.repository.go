package repository

import (
	"context"
	"hta-platform/internal/media/domain/model/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaRepository interface {
	GetMediaByUrl(ctx context.Context, url string) (entity.Media, error)

	GetMedias(ctx context.Context, req interface{}) ([]entity.Media, int64, error)

	// CreateMedia inserts a new media record. Uses ON CONFLICT DO NOTHING on the url column.
	// Returns the created media and a boolean indicating if it was actually inserted (false = duplicate).
	CreateMedia(tx *gorm.DB, media *entity.Media) (*entity.Media, bool, error)

	// UpdateMediaByURL updates a media record by its URL slug.
	// Returns the updated media, or nil if no media was found with the given URL.
	UpdateMediaByURL(tx *gorm.DB, url string, media *entity.Media) (*entity.Media, error)

	// InsertOtherName inserts a MediaOtherName record.
	InsertOtherName(tx *gorm.DB, otherName *entity.MediaOtherName) error

	// AttachAuthors bulk-inserts MediaToAuthor join records.
	AttachAuthors(tx *gorm.DB, mediaID uuid.UUID, authorIDs []uuid.UUID) error

	// AttachCategories bulk-inserts MediaToCategory join records.
	AttachCategories(tx *gorm.DB, mediaID uuid.UUID, categoryIDs []uuid.UUID) error

	// DeleteAuthorsByMediaID removes all author associations for a media.
	DeleteAuthorsByMediaID(tx *gorm.DB, mediaID uuid.UUID) error

	// DeleteCategoriesByMediaID removes all category associations for a media.
	DeleteCategoriesByMediaID(tx *gorm.DB, mediaID uuid.UUID) error

	// DeleteOtherNamesByMediaID removes all other names for a media.
	DeleteOtherNamesByMediaID(tx *gorm.DB, mediaID uuid.UUID) error
}
