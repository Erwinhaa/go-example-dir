package route

import (
	"myapp/internal/delivery/http/v1/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userController controller.UserController) {
	userRoutes := r.Group("users")
	{
		userRoutes.POST("", userController.CreateNewUser)
		userRoutes.GET("", userController.GetUser)
	}
}
