package usecase

import (
	"context"
)

type RoleUseCase struct {
	RoleRepogitory
}

func NewRoleUseCase(repogitory RoleRepogitory) *RoleUseCase {
	return &RoleUseCase{repogitory}
}

func (u RoleUseCase) CreateRole(ctx context.Context, roleName string) error {
	return u.RoleRepogitory.CreateRole(ctx, roleName)
}
