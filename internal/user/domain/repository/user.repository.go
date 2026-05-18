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
	GetBookmarkedAuthors(ctx context.Context, userID string) ([]BookmarkedAuthor, error)
	IsBookmarkedAuthor(ctx context.Context, userID string, authorID string) (bool, error)

	BookmarkMedia(ctx context.Context, userID string, mediaID string) error
	UnbookmarkMedia(ctx context.Context, userID string, mediaID string) error
	GetBookmarkedMedias(ctx context.Context, userID string) ([]mediaEntity.Media, error)
	IsBookmarkedMedia(ctx context.Context, userID string, mediaID string) (bool, error)

	UpsertReadingProgress(ctx context.Context, progress *entity.UserReadingProgress) error
	GetReadingProgress(ctx context.Context, userID string) ([]UserMediaProgress, error)

	StartReadingSession(ctx context.Context, session *entity.UserReadingSession) error
	EndReadingSession(ctx context.Context, sessionID string) error
	GetReadingSessions(ctx context.Context, userID string) ([]UserReadingSessionSummary, error)
}

type UserReadingSessionSummary struct {
	MediaID     string
	Duration    int64
	FirstReadAt string
	LastReadAt  string
}

type UserMediaProgress struct {
	mediaEntity.Media
	LastReadChapterOrder int64
	ChapterURL           string
	ChapterImageOrder    int
}

type BookmarkedAuthor struct {
	authorEntity.Author
	FirstMedia *mediaEntity.Media
}
