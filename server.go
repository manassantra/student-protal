package main

import (
	"github.com/gin-gonic/gin"
	"manassantra.online/golang-proj/student-portal/controller"
	"manassantra.online/golang-proj/student-portal/service"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func main() {
	server := gin.Default()

	server.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messege": "Hello User",
		})
	})

	server.GET("/finduser", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})
	server.GET("/finduser/:id", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindById(ctx))
	})
	server.POST("/createuser", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})

	server.Run(":8009")
}
