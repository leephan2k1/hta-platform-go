package repository

import (
	"context"
	"hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/media/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mediaChapterRepository struct {
	db *gorm.DB
}

// GetChapterImages implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) GetChapterImages(ctx context.Context, mediaUrl string, chapterUrl string) ([]entity.ChapterImage, error) {
	var chapter entity.MediaChapter
	err := m.db.WithContext(ctx).
		Joins("JOIN hta.media m ON m.id = hta.media_chapter.media_id").
		Where("m.url = ? AND hta.media_chapter.url = ?", mediaUrl, chapterUrl).
		Preload("Images", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"hta\".\"chapter_image\".\"order\" ASC")
		}).
		Preload("Images.Images").
		First(&chapter).Error

	if err != nil {
		return nil, err
	}

	return chapter.Images, nil
}

// FindChapterByMediaUrlAndOrder implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) FindChapterByMediaUrlAndOrder(ctx context.Context, mediaUrl string, order int64) (*entity.MediaChapter, error) {
	var chapter entity.MediaChapter
	err := m.db.WithContext(ctx).
		Joins("JOIN hta.media m ON m.id = hta.media_chapter.media_id").
		Where("m.url = ? AND hta.media_chapter.order = ?", mediaUrl, order).
		First(&chapter).Error

	if err != nil {
		return nil, err
	}

	return &chapter, nil
}

// CreateChapterImages implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) CreateChapterImages(ctx context.Context, images []*entity.ChapterImage) ([]entity.ChapterImage, error) {
	err := m.db.WithContext(ctx).Create(&images).Error
	if err != nil {
		return nil, err
	}

	res := make([]entity.ChapterImage, len(images))
	for i, img := range images {
		res[i] = *img
	}
	return res, nil
}

// CreateMediaChapters implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) CreateMediaChapters(ctx context.Context, chapters []*entity.MediaChapter) ([]entity.MediaChapter, error) {
	err := m.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "media_id"}, {Name: "url"}, {Name: "order"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "source", "language", "updated_at"}),
		}).Create(&chapters).Error

	if err != nil {
		return nil, err
	}

	res := make([]entity.MediaChapter, len(chapters))
	for i, ch := range chapters {
		res[i] = *ch
	}

	return res, nil
}

// GetMediaChaptersByMediaUrl implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) GetMediaChaptersByMediaUrl(ctx context.Context, url string) ([]entity.MediaChapter, error) {
	var chapters []entity.MediaChapter
	err := m.db.WithContext(ctx).
		Joins("JOIN hta.media m ON m.id = hta.media_chapter.media_id").
		Where("m.url = ?", url).
		Order("hta.media_chapter.order DESC").
		Find(&chapters).Error

	return chapters, err
}

func NewMediaChapterRepository(db *gorm.DB) repository.MediaChapterRepository {
	return &mediaChapterRepository{db: db}
}
