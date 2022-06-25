package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"manassantra.online/golang-proj/student-portal/entity"
)

type UserService interface {
	Save(*gin.Context)
	FindAll(*gin.Context)
	FindById(*gin.Context)
	// UpdateById() entity.UserEntity
	// DeleteById(*entity.UserEntity) entity.UserEntity
}

type userService struct {
	users []entity.UserEntity
}

func New() UserService {
	return &userService{}
}

func (service *userService) Save(c *gin.Context) {
	var user entity.UserEntity
	if err := c.BindJSON(&user); err != nil {
		return
	}
	service.users = append(service.users, user)
	c.IndentedJSON(http.StatusCreated, user)
	//return user
}

func (service *userService) FindAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.users)
}

func (service *userService) FindById(c *gin.Context) {
	//ID:= c.Param(`id`)
	//error := "User Not Exist"
	Id, _ := strconv.ParseInt(c.Param(`id`), 10, 64)
	for _, v := range service.users {
		if Id == v.Id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	//return error
	c.IndentedJSON(209, gin.H{"error": "No User Exist"})
}
