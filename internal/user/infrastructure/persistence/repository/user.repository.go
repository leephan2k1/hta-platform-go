package repository

import (
	"context"
	authorEntity "hta-platform/internal/author/domain/model/entity"
	mediaEntity "hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/user/domain/model/entity"
	"hta-platform/internal/user/domain/repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

// IsExistsUser implements [repository.UserRepository].
func (u *userRepository) IsExistsUser(ctx context.Context, id string) (bool, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

// CreateUser implements [repository.UserRepository].
func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

// BookmarkAuthor implements [repository.UserRepository].
func (u *userRepository) BookmarkAuthor(ctx context.Context, userID string, authorID string) error {
	aID, err := uuid.Parse(authorID)
	if err != nil {
		return err
	}
	bookmark := &entity.UserAuthor{
		UserID:   userID,
		AuthorID: aID,
	}
	return u.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(bookmark).Error
}

// UnbookmarkAuthor implements [repository.UserRepository].
func (u *userRepository) UnbookmarkAuthor(ctx context.Context, userID string, authorID string) error {
	return u.db.WithContext(ctx).Unscoped().
		Where("user_id = ? AND author_id = ?", userID, authorID).
		Delete(&entity.UserAuthor{}).Error
}

// GetBookmarkedAuthors implements [repository.UserRepository].
func (u *userRepository) GetBookmarkedAuthors(ctx context.Context, userID string) ([]repository.BookmarkedAuthor, error) {
	var authors []authorEntity.Author
	err := u.db.WithContext(ctx).
		Table("hta.author").
		Joins("JOIN hta.user_to_author ON hta.user_to_author.author_id = hta.author.id").
		Where("hta.user_to_author.user_id = ? AND hta.user_to_author.deleted_at IS NULL", userID).
		Find(&authors).Error
	if err != nil {
		return nil, err
	}

	res := make([]repository.BookmarkedAuthor, len(authors))
	for i, a := range authors {
		res[i].Author = a
		// Get first media info
		var media mediaEntity.Media
		err := u.db.WithContext(ctx).
			Table("hta.media").
			Joins("JOIN hta.media_to_author ON hta.media_to_author.media_id = hta.media.id").
			Where("hta.media_to_author.author_id = ? AND hta.media_to_author.deleted_at IS NULL", a.ID).
			Order("hta.media.created_at DESC").
			Preload("Images").
			First(&media).Error
		if err == nil {
			res[i].FirstMedia = &media
		}
	}
	return res, nil
}

// IsBookmarkedAuthor implements [repository.UserRepository].
func (u *userRepository) IsBookmarkedAuthor(ctx context.Context, userID string, authorID string) (bool, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&entity.UserAuthor{}).
		Where("user_id = ? AND author_id = ?", userID, authorID).
		Count(&count).Error
	return count > 0, err
}

// BookmarkMedia implements [repository.UserRepository].
func (u *userRepository) BookmarkMedia(ctx context.Context, userID string, mediaID string) error {
	mID, err := uuid.Parse(mediaID)
	if err != nil {
		return err
	}
	bookmark := &entity.UserMedia{
		UserID:  userID,
		MediaID: mID,
	}
	return u.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(bookmark).Error
}

// UnbookmarkMedia implements [repository.UserRepository].
func (u *userRepository) UnbookmarkMedia(ctx context.Context, userID string, mediaID string) error {
	return u.db.WithContext(ctx).Unscoped().
		Where("user_id = ? AND media_id = ?", userID, mediaID).
		Delete(&entity.UserMedia{}).Error
}

// GetBookmarkedMedias implements [repository.UserRepository].
func (u *userRepository) GetBookmarkedMedias(ctx context.Context, userID string) ([]mediaEntity.Media, error) {
	var medias []mediaEntity.Media
	err := u.db.WithContext(ctx).
		Table("hta.media").
		Joins("JOIN hta.user_to_media ON hta.user_to_media.media_id = hta.media.id").
		Where("hta.user_to_media.user_id = ? AND hta.user_to_media.deleted_at IS NULL", userID).
		Preload("Authors").
		Preload("Categories").
		Preload("Images").
		Preload("Chapters", func(db *gorm.DB) *gorm.DB {
			return db.Select("DISTINCT ON (media_id) *").Order("media_id, hta.media_chapter.order DESC")
		}).
		Find(&medias).Error
	return medias, err
}

// IsBookmarkedMedia implements [repository.UserRepository].
func (u *userRepository) IsBookmarkedMedia(ctx context.Context, userID string, mediaID string) (bool, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&entity.UserMedia{}).
		Where("user_id = ? AND media_id = ?", userID, mediaID).
		Count(&count).Error
	return count > 0, err
}

// UpsertReadingProgress implements [repository.UserRepository].
func (u *userRepository) UpsertReadingProgress(ctx context.Context, progress *entity.UserReadingProgress) error {
	return u.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_id"},
			{Name: "media_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"chapter_id", "chapter_image_order", "updated_at"}),
	}).Create(progress).Error
}

// GetReadingProgress implements [repository.UserRepository].
func (u *userRepository) GetReadingProgress(ctx context.Context, userID string) ([]repository.UserMediaProgress, error) {
	var progresses []struct {
		MediaID              uuid.UUID
		LastReadChapterOrder int64
	}

	err := u.db.WithContext(ctx).
		Table("hta.user_reading_progress p").
		Select("p.media_id, MAX(c.order) as last_read_chapter_order").
		Joins("JOIN hta.media_chapter c ON c.id = p.chapter_id").
		Where("p.user_id = ? AND p.deleted_at IS NULL", userID).
		Group("p.media_id").
		Find(&progresses).Error
	if err != nil {
		return nil, err
	}

	if len(progresses) == 0 {
		return nil, nil
	}

	res := make([]repository.UserMediaProgress, len(progresses))
	for i, p := range progresses {
		var media mediaEntity.Media
		err := u.db.WithContext(ctx).
			Preload("Authors").
			Preload("Categories").
			Preload("Chapters").
			Preload("Images").
			First(&media, "id = ?", p.MediaID).Error
		if err == nil {
			res[i].Media = media
			res[i].LastReadChapterOrder = p.LastReadChapterOrder
		}
	}
	return res, nil
}

// StartReadingSession implements [repository.UserRepository].
func (u *userRepository) StartReadingSession(ctx context.Context, session *entity.UserReadingSession) error {
	session.StartedAt = time.Now()
	return u.db.WithContext(ctx).Create(session).Error
}

// EndReadingSession implements [repository.UserRepository].
func (u *userRepository) EndReadingSession(ctx context.Context, sessionID string) error {
	var session entity.UserReadingSession
	if err := u.db.WithContext(ctx).First(&session, "id = ?", sessionID).Error; err != nil {
		return err
	}
	now := time.Now()
	session.EndedAt = now
	session.Duration = int64(now.Sub(session.StartedAt).Seconds())
	return u.db.WithContext(ctx).Save(&session).Error
}

// GetReadingSessions implements [repository.UserRepository].
func (u *userRepository) GetReadingSessions(ctx context.Context, userID string) ([]repository.UserReadingSessionSummary, error) {
	var results []repository.UserReadingSessionSummary
	err := u.db.WithContext(ctx).
		Table("hta.user_reading_sessions").
		Select("media_id, SUM(duration) as duration, TO_CHAR(MIN(started_at), 'YYYY-MM-DD') as first_read_at, TO_CHAR(MAX(started_at), 'YYYY-MM-DD') as last_read_at").
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Group("media_id").
		Scan(&results).Error
	return results, err
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}
