package controller

import (
	"heysay/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	RoleUseCase
}

func NewRoleController(u RoleUseCase) *RoleController {
	return &RoleController{u}
}

type GetRolesResponse struct {
	Roles []*entity.Role `json:"roles"`
}

// GetRoles godoc
//
// @Summary	全ロール取得 API
// @Description
// @Tags		Role
// @Accept		json
// @Produce		json
// @Success		200		{object}	GetRolesResponse		"OK"
// @Failure		401		{object}	entity.ErrorResponse
// @Failure		404		{object}	entity.ErrorResponse
// @Router		/roles	[get]
func (c RoleController) GeteRoles(ctx *gin.Context) (interface{}, error) {
	roles, err := c.RoleUseCase.GetRoles(ctx)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return GetRolesResponse{Roles: roles}, nil
}

type GetRoleResponse struct {
	Role *entity.Role `json:"role"`
}

// GetRoleByID godoc
//
// @Summary	ロール取得 API
// @Description
// @Tags		Role
// @Accept		json
// @Param		roleId			path		string					true	"ロール ID"
// @Produce		json
// @Success		200				{object}	GetRoleResponse			"OK"
// @Failure		401				{object}	entity.ErrorResponse
// @Failure		404				{object}	entity.ErrorResponse
// @Router		/roles/{roleId} [get]
func (c RoleController) GetRoleByID(ctx *gin.Context) (interface{}, error) {
	roleId := ctx.Param("roleId")
	role, err := c.RoleUseCase.GetRoleByID(ctx, roleId)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return GetRoleResponse{Role: role}, nil
}

type CreateRoleRequest struct {
	Name string `form:"name"`
}

// CreateRole godoc
//
// @Summary	ロール作成 API
// @Description
// @Tags		Role
// @Accept		json
// @Produce		json
// @Param		request	body		CreateRoleRequest		true	"ロール作成リクエスト"
// @Success		201		"Created"
// @Failure		400		{object}	entity.ErrorResponse
// @Failure		401		{object}	entity.ErrorResponse
// @Router		/roles [post]
func (c RoleController) CreateRole(ctx *gin.Context) (interface{}, error) {
	var req *CreateRoleRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return nil, c.RoleUseCase.CreateRole(ctx, req.Name)
}
