package controller

import (
	"myapp/internal/dto"
	"myapp/internal/usecase"
	"myapp/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	UserController interface {
		CreateNewUser(ctx *gin.Context)
		GetUser(ctx *gin.Context)
		GetUserList(ctx *gin.Context)
	}

	userController struct {
		uc usecase.UserUseCase
	}
)

func NewUserController(uc usecase.UserUseCase) *userController {
	return &userController{uc}
}

var _ UserController = (*userController)(nil)

func (c *userController) CreateNewUser(ctx *gin.Context) {
	var payload *dto.CreateUserRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Kontol")
		return
	}

	newUser, err := c.uc.CreateUser(ctx, *payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Kontol")
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}

func (c *userController) GetUser(ctx *gin.Context) {
	var userId dto.GetUserRequest
	if err := ctx.ShouldBindUri(&userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Kontol")
		return
	}

	user, err := c.uc.GetUserById(ctx, userId.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res := utils.NewFailedResponse()
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Goblok")
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "ANjing")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *userController) GetUserList(ctx *gin.Context) {

}
