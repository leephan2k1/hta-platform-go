package service

import (
	"context"

	"hta-platform/internal/user/domain/model/entity"
	userRepo "hta-platform/internal/user/domain/repository"
)

type userService struct {
	userRepo userRepo.UserRepository
}

// IsExistsUser implements UserService.
func (as *userService) IsExistsUser(ctx context.Context, id int64) (*entity.Account, error) {
	panic("unimplemented")
}

func NewUserService(userRepo userRepo.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
