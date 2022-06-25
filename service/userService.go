package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"manassantra.online/golang-proj/student-portal/entity"
)

type UserService interface {
	Save(entity.UserEntity) entity.UserEntity
	FindAll() []entity.UserEntity
	FindById(*gin.Context) any
	// UpdateById() entity.UserEntity
	// DeleteById(*entity.UserEntity) entity.UserEntity
}

type userService struct {
	users []entity.UserEntity
}

func New() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.UserEntity) entity.UserEntity {
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindAll() []entity.UserEntity {
	return service.users
}

func (service *userService) FindById(c *gin.Context) any {
	//ID:= c.Param(`id`)
	error := "User Not Exist"
	Id, _ := strconv.ParseInt(c.Param(`id`), 10, 64)
	for _, v := range service.users {
		if Id == v.Id {
			return v
		}
	}
	return error
	// c.IndentedJSON(209, gin.H{"error": "No User Exist"})
}
