package repository

import (
	"context"
	"hta-platform/internal/media/controller/dto"
	"hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/media/domain/repository"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mediaRepository struct {
	db *gorm.DB
}

// GetMediaByUrl implements [repository.MediaRepository].
func (m *mediaRepository) GetMediaByUrl(ctx context.Context, url string) (entity.Media, error) {
	var media entity.Media
	err := m.db.WithContext(ctx).
		Preload("Status").
		Preload("Type").
		Preload("Categories").
		Preload("Authors").
		Preload("OtherNames").
		Preload("Chapters", func(db *gorm.DB) *gorm.DB {
			return db.Order("hta.media_chapter.order DESC")
		}).
		Where("url = ?", url).
		First(&media).Error

	return media, err
}

// GetMedias implements [repository.MediaRepository].
func (m *mediaRepository) GetMedias(ctx context.Context, req interface{}) ([]entity.Media, int64, error) {
	r, ok := req.(*dto.GetMediasReq)
	if !ok {
		return nil, 0, nil
	}

	query := m.db.WithContext(ctx).Model(&entity.Media{})

	// 1. Filtering
	if r.Name != "" {
		query = query.Where("LOWER(name) ILIKE ?", "%"+strings.ToLower(r.Name)+"%")
	}

	query = query.Where("is_nsfw = ?", r.IsNSFW)
	if r.SysStatus != "" {
		query = query.Where("sys_status = ?", r.SysStatus)
	}

	if len(r.Authors) > 0 {
		query = query.Joins("JOIN hta.media_to_author ma ON ma.media_id = hta.media.id").
			Joins("JOIN hta.author a ON a.id = ma.author_id").
			Where("a.author_url IN ?", r.Authors)
	}

	if len(r.Categories) > 0 {
		var includeSlugs []string
		var excludeSlugs []string
		for _, cat := range r.Categories {
			if strings.HasPrefix(cat, "!") {
				excludeSlugs = append(excludeSlugs, strings.TrimPrefix(cat, "!"))
			} else {
				includeSlugs = append(includeSlugs, cat)
			}
		}

		if len(includeSlugs) > 0 {
			query = query.Joins("JOIN hta.media_to_category mc ON mc.media_id = hta.media.id").
				Joins("JOIN hta.category c ON c.id = mc.category_id").
				Where("c.slug IN ?", includeSlugs)
		}

		if len(excludeSlugs) > 0 {
			query = query.Where("NOT EXISTS (SELECT 1 FROM hta.media_to_category mc_ex JOIN hta.category c_ex ON c_ex.id = mc_ex.category_id WHERE mc_ex.media_id = hta.media.id AND c_ex.slug IN ?)", excludeSlugs)
		}
	}

	// 2. Count total before pagination
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 3. Preload for info
	query = query.Preload("Authors").Preload("Categories").
		Preload("Chapters", func(db *gorm.DB) *gorm.DB {
			return db.Select("DISTINCT ON (media_id) *").Order("media_id, hta.media_chapter.order DESC")
		})

	// 4. Sorting
	if len(r.SortBy) > 0 {
		for _, sort := range r.SortBy {
			switch strings.ToLower(sort) {
			case "updatedat":
				query = query.Order("updated_at DESC")
			case "name":
				query = query.Order("name DESC")
			}
		}
	} else {
		query = query.Order("updated_at DESC")
	}

	// 5. Pagination
	limit := r.Limit
	if limit <= 0 {
		limit = 25
	}
	page := r.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	var items []entity.Media
	if err := query.Limit(limit).Offset(offset).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
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
