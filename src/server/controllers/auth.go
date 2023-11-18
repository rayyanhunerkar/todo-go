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

// Register godoc
// @Summary Register
// @Description Register a new User
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.RegisterRequest true "RegisterRequest"
// @Success 200 {object} models.Response
// @Router /auth/signup [post]
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

// Login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "LoginRequest"
// @Success 200 {object} models.Response
// @Router /auth/login [post]
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

// Me godoc
// @Summary Me
// @Description Gets current user
// @Tags User
// @Security Bearer
// @Produce json
// @Success 200 {object} models.Response
// @Router /user/me [get]
func (h *UserController) Me(context *gin.Context) {

	var response models.Response
	var userResponse models.MeResponse

	user, err := h.service.GetUserByID(string(context.GetString("currentUser")))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, "Wrong Username/Password")
		panic(err)
	}
	userResponse.ID = user.ID
	userResponse.Username = user.Username
	userResponse.FirstName = user.FirstName
	userResponse.LastName = user.LastName

	response.Data = userResponse
	response.Message = "User Fetched Successfully"

	context.JSON(http.StatusOK, response)
}
