package controller

import (
	"heysay/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	RoleUseCase
}

func NewRoleController(u RoleUseCase) *RoleController {
	return &RoleController{u}
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
// @Success		200		"OK"
// @Failure		400		{object}	entity.ErrorResponse
// @Failure		401		{object}	entity.ErrorResponse
// @Router		/roles [post]
func (c RoleController) CreateRole(ctx *gin.Context) (interface{}, error) {
	log.Print("lsdfghlsdgjlsdglj")
	var req *CreateRoleRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return nil, c.RoleUseCase.CreateRole(ctx, req.Name)
}
