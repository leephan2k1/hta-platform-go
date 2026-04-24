package repository

import (
	"context"

	"hta-platform/internal/user/domain/model/entity"
)

type UserRepository interface {
	// Get account by ID
	IsExistsUser(ctx context.Context, accountId int64) (*entity.Account, error)
}
