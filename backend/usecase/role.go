package usecase

import (
	"context"
	"heysay/entity"
)

type RoleUseCase struct {
	RoleRepogitory
}

func NewRoleUseCase(repogitory RoleRepogitory) *RoleUseCase {
	return &RoleUseCase{repogitory}
}

func (u RoleUseCase) GetRoles(ctx context.Context) ([]*entity.Role, error) {
	return u.RoleRepogitory.GetRoles(ctx)
}

func (u RoleUseCase) CreateRole(ctx context.Context, roleName string) error {
	return u.RoleRepogitory.CreateRole(ctx, roleName)
}
