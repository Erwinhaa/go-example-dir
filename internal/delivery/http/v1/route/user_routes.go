package route

import (
	"myapp/internal/delivery/http/v1/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, c controller.UserController) {
	r.POST("", c.CreateNewUser)
	r.GET("/:id", c.GetUser)
}
