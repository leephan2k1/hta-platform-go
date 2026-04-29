package repository

import (
	"context"
	"hta-platform/internal/user/domain/model/entity"
	"hta-platform/internal/user/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// IsExistsUser implements [repository.UserRepository].
func (u *userRepository) IsExistsUser(ctx context.Context, id string) (bool, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

// CreateUser implements [repository.UserRepository].
func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}
