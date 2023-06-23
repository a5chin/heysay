package controller

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock.go -package=$GOPACKAGE

import (
	"context"
	"heysay/entity"
)

type UserUseCase interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}
