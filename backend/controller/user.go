package controller

import (
	"heysay/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase
}

func NewUserController(u UserUseCase) *UserController {
	return &UserController{u}
}

type GetUsersResponse struct {
	Users []*entity.User `json:"users"`
}

// GetUsers godoc
//
// @Summary	全ユーザー取得 API
// @Description
// @Tags		User
// @Accept		json
// @Produce		json
// @Success		200		{object}	GetUsersResponse	"OK"
// @Router		/users	[get]
func (c UserController) GetUsers(ctx *gin.Context) (interface{}, error) {
	users, err := c.UserUseCase.GetUsers(ctx)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return GetUsersResponse{Users: users}, nil
}

type GetUserResponse struct {
	User *entity.User `json:"user"`
}

// GetUserByID godoc
//
// @Summary	ユーザー取得 API
// @Description
// @Tags		User
// @Accept		json
// @Param		userId			path		string					true	"ユーザーID"
// @Produce		json
// @Success		200				{object}	GetUserResponse			"OK"
// @Failure		401				{object}	entity.ErrorResponse
// @Failure		404				{object}	entity.ErrorResponse
// @Router		/users/{userId} [get]
func (c UserController) GetUserByID(ctx *gin.Context) (interface{}, error) {
	userID := ctx.Param("userId")
	user, err := c.UserUseCase.GetUserByID(ctx, userID)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return GetUserResponse{User: user}, nil
}
