package service

import (
	"context"
	authorDto "hta-platform/internal/author/controller/dto"
	mediaDto "hta-platform/internal/media/controller/dto"
	"hta-platform/internal/user/controller/dto"
	"hta-platform/internal/user/domain/model/entity"
	"hta-platform/internal/user/domain/repository"
	"github.com/google/uuid"
)

type userService struct {
	userRepo repository.UserRepository
}

// RegisterUser implements [UserService].
func (u *userService) RegisterUser(ctx context.Context, req dto.RegisterUserReq) error {
	exists, err := u.userRepo.IsExistsUser(ctx, req.Auth0Id)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	user := &entity.User{
		ID:        req.Auth0Id,
		FirstName: req.GivenName,
		LastName:  req.FamilyName,
		Email:     req.Email,
		Picture:   req.Picture,
	}

	return u.userRepo.CreateUser(ctx, user)
}

// BookmarkAuthor implements [UserService].
func (u *userService) BookmarkAuthor(ctx context.Context, userID string, authorID string) error {
	return u.userRepo.BookmarkAuthor(ctx, userID, authorID)
}

// UnbookmarkAuthor implements [UserService].
func (u *userService) UnbookmarkAuthor(ctx context.Context, userID string, authorID string) error {
	return u.userRepo.UnbookmarkAuthor(ctx, userID, authorID)
}

// GetBookmarkedAuthors implements [UserService].
func (u *userService) GetBookmarkedAuthors(ctx context.Context, userID string) ([]authorDto.AuthorRes, error) {
	authors, err := u.userRepo.GetBookmarkedAuthors(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]authorDto.AuthorRes, len(authors))
	for i, a := range authors {
		res[i].SetData(a.Author)
		if a.FirstMedia != nil {
			res[i].FirstMedia = &authorDto.AuthorMediaRes{}
			res[i].FirstMedia.SetData(*a.FirstMedia)
		}
	}
	return res, nil
}

// IsBookmarkedAuthor implements [UserService].
func (u *userService) IsBookmarkedAuthor(ctx context.Context, userID string, authorID string) (bool, error) {
	return u.userRepo.IsBookmarkedAuthor(ctx, userID, authorID)
}

// BookmarkMedia implements [UserService].
func (u *userService) BookmarkMedia(ctx context.Context, userID string, mediaID string) error {
	return u.userRepo.BookmarkMedia(ctx, userID, mediaID)
}

// UnbookmarkMedia implements [UserService].
func (u *userService) UnbookmarkMedia(ctx context.Context, userID string, mediaID string) error {
	return u.userRepo.UnbookmarkMedia(ctx, userID, mediaID)
}

// GetBookmarkedMedias implements [UserService].
func (u *userService) GetBookmarkedMedias(ctx context.Context, userID string) ([]mediaDto.MediaResponse, error) {
	medias, err := u.userRepo.GetBookmarkedMedias(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]mediaDto.MediaResponse, len(medias))
	for i, media := range medias {
		res[i].SetData(media)
	}
	return res, nil
}

// IsBookmarkedMedia implements [UserService].
func (u *userService) IsBookmarkedMedia(ctx context.Context, userID string, mediaID string) (bool, error) {
	return u.userRepo.IsBookmarkedMedia(ctx, userID, mediaID)
}

// UpsertReadingProgress implements [UserService].
func (u *userService) UpsertReadingProgress(ctx context.Context, userID string, req dto.UserReadingProgressReq) error {
	cID, err := uuid.Parse(req.ChapterID)
	if err != nil {
		return err
	}
	mID, err := uuid.Parse(req.MediaID)
	if err != nil {
		return err
	}

	progress := &entity.UserReadingProgress{
		UserID:            userID,
		ChapterID:         cID,
		MediaID:           mID,
		ChapterImageOrder: *req.ImageOrder,
	}

	return u.userRepo.UpsertReadingProgress(ctx, progress)
}

// GetReadingProgress implements [UserService].
func (u *userService) GetReadingProgress(ctx context.Context, userID string) ([]mediaDto.MediaResponse, error) {
	progresses, err := u.userRepo.GetReadingProgress(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]mediaDto.MediaResponse, len(progresses))
	for i, p := range progresses {
		res[i].SetData(p.Media)

		// Find max chapter order
		var maxOrder int64
		for _, ch := range p.Media.Chapters {
			if ch.Order > maxOrder {
				maxOrder = ch.Order
			}
		}

		if maxOrder > 0 {
			res[i].Progress = float64(p.LastReadChapterOrder) / float64(maxOrder)
		} else {
			res[i].Progress = 0
		}
		res[i].ChapterProgress = p.ChapterURL
		res[i].ChapterImageProgress = p.ChapterImageOrder
	}
	return res, nil
}

// StartReadingSession implements [UserService].
func (u *userService) StartReadingSession(ctx context.Context, userID string, req dto.UserReadingSessionReq) (*dto.UserReadingSessionStartRes, error) {
	mID, err := uuid.Parse(req.MediaID)
	if err != nil {
		return nil, err
	}

	session := &entity.UserReadingSession{
		UserID:  userID,
		MediaID: mID,
	}

	if err := u.userRepo.StartReadingSession(ctx, session); err != nil {
		return nil, err
	}

	return &dto.UserReadingSessionStartRes{
		SessionID: session.ID.String(),
	}, nil
}

// EndReadingSession implements [UserService].
func (u *userService) EndReadingSession(ctx context.Context, userID string, req dto.UserReadingSessionEndReq) error {
	return u.userRepo.EndReadingSession(ctx, req.SessionID)
}

// GetReadingSessions implements [UserService].
func (u *userService) GetReadingSessions(ctx context.Context, userID string) (map[string]dto.UserReadingSessionRes, error) {
	summaries, err := u.userRepo.GetReadingSessions(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make(map[string]dto.UserReadingSessionRes)
	for _, s := range summaries {
		res[s.MediaID] = dto.UserReadingSessionRes{
			MediaID:     s.MediaID,
			Duration:    s.Duration,
			FirstReadAt: s.FirstReadAt,
			LastReadAt:  s.LastReadAt,
		}
	}
	return res, nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
