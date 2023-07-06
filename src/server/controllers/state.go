package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/todo-go/src/database/models"
	"github.com/rayyanhunerkar/todo-go/src/database/repositories"
)

type StateController struct {
	service models.StateRepository
}

func InitStateController(stateRepo *repositories.StateRepo) *StateController {
	return &StateController{
		service: stateRepo,
	}
}

func (h *StateController) CreateState(context *gin.Context) {

	var request models.StateCreateRequest
	if err := context.BindJSON(&request); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	state, err := h.service.CreateState(request)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusCreated, &state)
}

func (h *StateController) GetStates(context *gin.Context) {
	states, err := h.service.GetStates()
	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}
	context.JSON(http.StatusOK, &states)
}

func (h *StateController) GetStateByID(context *gin.Context) {
	params := context.Param("id")
	state, err := h.service.GetStateByID(params)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, "State not found")
	}
	context.JSON(http.StatusOK, state)
}
