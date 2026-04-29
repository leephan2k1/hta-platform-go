package repository

import (
	"context"
	"hta-platform/internal/user/domain/model/entity"
)

type UserRepository interface {
	IsExistsUser(ctx context.Context, id string) (bool, error)

	CreateUser(ctx context.Context, user *entity.User) error
}
