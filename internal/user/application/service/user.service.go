package service

import (
	"context"

	"hta-platform/internal/user/domain/model/entity"
)

type UserService interface {
	IsExistsUser(ctx context.Context, id int64) (*entity.Account, error)
}
