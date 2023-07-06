package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/todo-go/src/database/models"
	"github.com/rayyanhunerkar/todo-go/src/database/repositories"
	"github.com/spf13/viper"
)

type UserController struct {
	service models.UserRepository
	conf    *viper.Viper
}

func InitUserController(userRepo *repositories.UserRepo, conf *viper.Viper) *UserController {
	return &UserController{
		service: userRepo,
		conf:    conf,
	}
}

func (h *UserController) Register(context *gin.Context) {
	var request models.RegisterRequest

	if err := context.BindJSON(&request); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.service.CreateUser(request)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, &user)
}

func (h *UserController) Login(context *gin.Context) {
	var request models.LoginRequest

	if err := context.BindJSON(&request); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Login(request, h.conf)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, "Wrong Username/Password")
		panic(err)
	}
	context.JSON(http.StatusOK, token)
}
