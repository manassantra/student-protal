package controller

import (
	"github.com/gin-gonic/gin"
	"manassantra.online/golang-proj/student-portal/entity"
	"manassantra.online/golang-proj/student-portal/service"
)

type UserController interface {
	FindAll() []entity.UserEntity
	Save(*gin.Context) entity.UserEntity
	FindById(*gin.Context) any
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.UserEntity {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.UserEntity {
	var user entity.UserEntity
	ctx.BindJSON(&user)
	c.service.Save(user)
	return user
}

func (c *controller) FindById(ctx *gin.Context) any {
	// c.service.FindById(ctx)
	return c.service.FindById(ctx)
}
