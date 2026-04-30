package service

import (
	"context"
	authorDto "hta-platform/internal/author/controller/dto"
	mediaDto "hta-platform/internal/media/controller/dto"
	"hta-platform/internal/user/controller/dto"
	"hta-platform/internal/user/domain/model/entity"
	"hta-platform/internal/user/domain/repository"
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

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
