package repository

import (
	"context"
	authorEntity "hta-platform/internal/author/domain/model/entity"
	mediaEntity "hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/user/domain/model/entity"
)

type UserRepository interface {
	IsExistsUser(ctx context.Context, id string) (bool, error)

	CreateUser(ctx context.Context, user *entity.User) error
	
	BookmarkAuthor(ctx context.Context, userID string, authorID string) error
	UnbookmarkAuthor(ctx context.Context, userID string, authorID string) error
	GetBookmarkedAuthors(ctx context.Context, userID string) ([]authorEntity.Author, error)

	BookmarkMedia(ctx context.Context, userID string, mediaID string) error
	UnbookmarkMedia(ctx context.Context, userID string, mediaID string) error
	GetBookmarkedMedias(ctx context.Context, userID string) ([]mediaEntity.Media, error)
}
