package router

import (
	"restapi/controllers"
	"restapi/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRouter(r *gin.Engine) {
	PhotoRouter := r.Group("/photos")
	PhotoRouter.POST("/", middlewares.Authentication(), controllers.CreatePhoto)
	PhotoRouter.GET("/", controllers.GetPhotos)
	PhotoRouter.GET("/:photo_id", controllers.GetPhoto)
	PhotoRouter.PUT("/:photo_id", middlewares.Authentication(), controllers.UpdatePhoto)
	PhotoRouter.DELETE("/:photo_id", middlewares.Authentication(), controllers.DeletePhoto)
}
