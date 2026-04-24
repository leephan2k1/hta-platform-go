package repository

import (
	"context"
	"log"

	"hta-platform/internal/user/domain/model/entity"
	"hta-platform/internal/user/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// IsExistsUser implements repository.UserRepository.
func (u *userRepository) IsExistsUser(ctx context.Context, accountId int64) (*entity.Account, error) {
	var account entity.Account
	log.Println("Checking if user exists with ID:", accountId)
	return &account, nil
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}
