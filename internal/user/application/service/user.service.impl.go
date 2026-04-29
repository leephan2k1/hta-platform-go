package service

import (
	"context"
	"hta-platform/internal/user/controller/dto"
	"hta-platform/internal/user/domain/model/entity"
	userRepo "hta-platform/internal/user/domain/repository"
)

type userService struct {
	userRepo userRepo.UserRepository
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
		Email:     req.Email,
		FirstName: req.GivenName,
		LastName:  req.FamilyName,
		Picture:   req.Picture,
	}

	return u.userRepo.CreateUser(ctx, user)
}

func NewUserService(userRepo userRepo.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
