package usecase

import (
	"context"
	"heysay/entity"
)

type UserUseCase struct {
	UserRepogitory
}

func NewUserUseCase(repogitory UserRepogitory) *UserUseCase {
	return &UserUseCase{repogitory}
}

func (u UserUseCase) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return u.UserRepogitory.GetUsers(ctx)
}

func (u UserUseCase) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	return u.UserRepogitory.GetUserByID(ctx, userID)
}
