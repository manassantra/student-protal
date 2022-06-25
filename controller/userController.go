package controller

import (
	"github.com/gin-gonic/gin"
	"manassantra.online/golang-proj/student-portal/service"
)

type UserController interface {
	FindAll(*gin.Context)
	Save(*gin.Context)
	FindById(*gin.Context)
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context) {
	c.service.FindAll(ctx)
}

func (c *controller) Save(ctx *gin.Context) {
	c.service.Save(ctx)
}

func (c *controller) FindById(ctx *gin.Context) {
	//var user entity.UserEntity
	c.service.FindById(ctx)

}
