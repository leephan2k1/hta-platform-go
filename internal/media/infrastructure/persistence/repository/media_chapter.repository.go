package repository

import (
	"context"
	"hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/media/domain/repository"

	"gorm.io/gorm"
)

type mediaChapterRepository struct {
	db *gorm.DB
}

// GetChapterImagesByChapterUrl implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) GetChapterImagesByChapterUrl(ctx context.Context, url string) ([]entity.ChapterImage, error) {
	var images []entity.ChapterImage
	err := m.db.WithContext(ctx).
		Joins("JOIN hta.media_chapter mc ON mc.id = hta.chapter_image.chapter_id").
		Where("mc.url = ?", url).
		Order("hta.chapter_image.order ASC").
		Find(&images).Error
	return images, err
}

// GetMediaChaptersByMediaUrl implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) GetMediaChaptersByMediaUrl(ctx context.Context, url string) ([]entity.MediaChapter, error) {
	var chapters []entity.MediaChapter
	err := m.db.WithContext(ctx).
		Joins("JOIN hta.media m ON m.id = hta.media_chapter.media_id").
		Where("m.url = ?", url).
		Order("hta.media_chapter.order ASC").
		Find(&chapters).Error
	return chapters, err
}

func NewMediaChapterRepository(db *gorm.DB) repository.MediaChapterRepository {
	return &mediaChapterRepository{db: db}
}
