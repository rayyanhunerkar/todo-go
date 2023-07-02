package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/todo-go/src/database/models"
	"github.com/rayyanhunerkar/todo-go/src/database/repositories"
)

type UserController struct {
	service models.UserRepository
}

func InitController(userRepo *repositories.UserRepo) *UserController {
	return &UserController{
		service: userRepo,
	}
}

func (h *UserController) register(r *gin.Context) {
	var request models.RegisterRequest

	if err := r.BindJSON(&request); err != nil {
		r.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.service.CreateUser(request)
	if err != nil {
		r.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	r.JSON(http.StatusCreated, &user)
}
