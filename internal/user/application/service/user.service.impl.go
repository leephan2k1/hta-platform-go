package service

import (
	userRepo "hta-platform/internal/user/domain/repository"
)

type userService struct {
	userRepo userRepo.UserRepository
}

func NewUserService(userRepo userRepo.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
