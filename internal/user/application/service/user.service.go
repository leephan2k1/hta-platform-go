package service

import (
	"context"
	"hta-platform/internal/user/controller/dto"
)

type UserService interface {
	RegisterUser(ctx context.Context, req dto.RegisterUserReq) error
}
