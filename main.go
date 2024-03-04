package main

import (
	"restapi/database"
	"restapi/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.UserRouter(r)
	router.PhotoRouter(r)
	database.ConnectDatabase()
	r.Run(":8080")
}
