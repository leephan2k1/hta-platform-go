package service

import (
	"context"

	"github.com/leedev/go-rest-ddd/internal/user/domain/model/entity"
)

type UserService interface {
	IsExistsUser(ctx context.Context, id int64) (*entity.Account, error)
}
