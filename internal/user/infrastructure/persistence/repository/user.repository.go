package repository

import (
	"context"
	authorEntity "hta-platform/internal/author/domain/model/entity"
	mediaEntity "hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/user/domain/model/entity"
	"hta-platform/internal/user/domain/repository"

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

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}
