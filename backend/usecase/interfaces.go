package usecase

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock.go -package=$GOPACKAGE

import (
	"context"
	"heysay/entity"
)

type UserRepogitory interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}

type RoleRepogitory interface {
	GetRoles(ctx context.Context) ([]*entity.Role, error)
	CreateRole(ctx context.Context, roleName string) error
}
