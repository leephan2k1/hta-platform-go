package repository

import (
	"hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/media/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mediaRepository struct {
	db *gorm.DB
}

// CreateMedia implements [repository.MediaRepository].
// Uses ON CONFLICT DO NOTHING on the url column. Returns false if the row was a duplicate.
func (m *mediaRepository) CreateMedia(tx *gorm.DB, media *entity.Media) (*entity.Media, bool, error) {
	result := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "url"}},
		DoNothing: true,
	}).Create(media)

	if result.Error != nil {
		return nil, false, result.Error
	}

	// RowsAffected == 0 means ON CONFLICT DO NOTHING triggered (duplicate slug/url)
	if result.RowsAffected == 0 {
		return nil, false, nil
	}

	return media, true, nil
}

// UpdateMediaByURL implements [repository.MediaRepository].
// Updates media fields WHERE url = ? and returns the updated record.
func (m *mediaRepository) UpdateMediaByURL(tx *gorm.DB, url string, media *entity.Media) (*entity.Media, error) {
	// Find the existing media by URL first
	var existing entity.Media
	if err := tx.Where("url = ?", url).First(&existing).Error; err != nil {
		return nil, err
	}

	// Update the fields on the existing record
	existing.Name = media.Name
	existing.Description = media.Description
	existing.URL = media.URL
	existing.StatusID = media.StatusID
	existing.TypeID = media.TypeID
	existing.IsNSFW = media.IsNSFW
	existing.Thumbnail = media.Thumbnail

	if err := tx.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

// InsertOtherName implements [repository.MediaRepository].
func (m *mediaRepository) InsertOtherName(tx *gorm.DB, otherName *entity.MediaOtherName) error {
	return tx.Create(otherName).Error
}

// AttachAuthors implements [repository.MediaRepository].
func (m *mediaRepository) AttachAuthors(tx *gorm.DB, mediaID uuid.UUID, authorIDs []uuid.UUID) error {
	records := make([]entity.MediaToAuthor, len(authorIDs))
	for i, authorID := range authorIDs {
		records[i] = entity.MediaToAuthor{
			MediaID:  mediaID,
			AuthorID: authorID,
		}
	}
	return tx.Create(&records).Error
}

// AttachCategories implements [repository.MediaRepository].
func (m *mediaRepository) AttachCategories(tx *gorm.DB, mediaID uuid.UUID, categoryIDs []uuid.UUID) error {
	records := make([]entity.MediaToCategory, len(categoryIDs))
	for i, categoryID := range categoryIDs {
		records[i] = entity.MediaToCategory{
			MediaID:    mediaID,
			CategoryID: categoryID,
		}
	}
	return tx.Create(&records).Error
}

// DeleteAuthorsByMediaID implements [repository.MediaRepository].
func (m *mediaRepository) DeleteAuthorsByMediaID(tx *gorm.DB, mediaID uuid.UUID) error {
	return tx.Where("media_id = ?", mediaID).Delete(&entity.MediaToAuthor{}).Error
}

// DeleteCategoriesByMediaID implements [repository.MediaRepository].
func (m *mediaRepository) DeleteCategoriesByMediaID(tx *gorm.DB, mediaID uuid.UUID) error {
	return tx.Where("media_id = ?", mediaID).Delete(&entity.MediaToCategory{}).Error
}

// DeleteOtherNamesByMediaID implements [repository.MediaRepository].
func (m *mediaRepository) DeleteOtherNamesByMediaID(tx *gorm.DB, mediaID uuid.UUID) error {
	return tx.Where("media_id = ?", mediaID).Delete(&entity.MediaOtherName{}).Error
}

func NewMediaRepository(db *gorm.DB) repository.MediaRepository {
	return &mediaRepository{db: db}
}
