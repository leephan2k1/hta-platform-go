package service

import (
	"context"

	"github.com/leedev/go-rest-ddd/internal/user/domain/model/entity"
	userRepo "github.com/leedev/go-rest-ddd/internal/user/domain/repository"
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
