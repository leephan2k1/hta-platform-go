package service

import (
	"context"
	authorDto "hta-platform/internal/author/controller/dto"
	mediaDto "hta-platform/internal/media/controller/dto"
	"hta-platform/internal/user/controller/dto"
)

type UserService interface {
	RegisterUser(ctx context.Context, req dto.RegisterUserReq) error

	BookmarkAuthor(ctx context.Context, userID string, authorID string) error
	UnbookmarkAuthor(ctx context.Context, userID string, authorID string) error
	GetBookmarkedAuthors(ctx context.Context, userID string) ([]authorDto.AuthorRes, error)
	IsBookmarkedAuthor(ctx context.Context, userID string, authorID string) (bool, error)

	BookmarkMedia(ctx context.Context, userID string, mediaID string) error
	UnbookmarkMedia(ctx context.Context, userID string, mediaID string) error
	GetBookmarkedMedias(ctx context.Context, userID string) ([]mediaDto.MediaResponse, error)
	IsBookmarkedMedia(ctx context.Context, userID string, mediaID string) (bool, error)

	UpsertReadingProgress(ctx context.Context, userID string, req dto.UserReadingProgressReq) error
	GetReadingProgress(ctx context.Context, userID string) ([]mediaDto.MediaResponse, error)
}
