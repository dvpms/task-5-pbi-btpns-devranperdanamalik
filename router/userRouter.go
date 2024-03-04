package router

import (
	"restapi/controllers"
	"restapi/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	UserRouter := r.Group("/users")
	UserRouter.POST("/register", controllers.CreateUser)
	UserRouter.POST("/login", controllers.Login)
	UserRouter.PUT("/:user_id", middlewares.Authentication(), controllers.UpdateUser)
	UserRouter.DELETE("/:user_id", middlewares.Authentication(), controllers.DeleteUser)
}
