package repository

import (
	"context"

	"github.com/leedev/go-rest-ddd/internal/user/domain/model/entity"
)

type UserRepository interface {
	// Get account by ID
	IsExistsUser(ctx context.Context, accountId int64) (*entity.Account, error)
}
