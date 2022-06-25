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

	server.GET("/finduser", userController.FindAll)
	server.GET("/finduser/:id", userController.FindById)
	server.POST("/createuser", userController.Save)

	server.Run(":8009")
}
