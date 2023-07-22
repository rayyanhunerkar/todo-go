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

// States godoc
// @Summary Create State
// @Security Bearer
// @Description Create a new state
// @Tags State
// @Accept json
// @Produce json
// @Param state body models.StateCreateRequest true "StateRequest"
// @Success 201 {object} models.Response
// @Router /states [post]
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

// States godoc
// @Summary Create State
// @Security Bearer
// @Description Create a new state
// @Tags State
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /states [get]
func (h *StateController) GetStates(context *gin.Context) {
	states, err := h.service.GetStates()
	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}
	context.JSON(http.StatusOK, &states)
}

// States godoc
// @Summary Create State
// @Security Bearer
// @Description Create a new state
// @Tags State
// @Accept json
// @Produce json
// @Param id path string true "fetch by id"
// @Success 200 {object} models.Response
// @Router /states/{id} [post]
func (h *StateController) GetStateByID(context *gin.Context) {
	params := context.Param("id")
	state, err := h.service.GetStateByID(params)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, "State not found")
	}
	context.JSON(http.StatusOK, state)
}

// States godoc
// @Summary Create State
// @Security Bearer
// @Description Create a new state
// @Tags State
// @Accept json
// @Produce json
// @Param id path string true "fetch by id"
// @Param state body models.StateUpdateRequest true "StateRequest"
// @Success 200 {object} models.Response
// @Router /states/{id} [patch]
func (h *StateController) UpdateState(context *gin.Context) {
	var request models.StateUpdateRequest
	id := context.Param("id")
	if err := context.BindJSON(&request); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, err := h.service.UpdateState(request, id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, "State not found")
	}
	context.JSON(http.StatusOK, response)
}

// States godoc
// @Summary Create State
// @Security Bearer
// @Description Create a new state
// @Tags State
// @Accept json
// @Produce json
// @Param id path string true "fetch by id"
// @Success 204
// @Router /states/{id} [delete]
func (h *StateController) DeleteState(context *gin.Context) {
	id := context.Param("id")
	err := h.service.DeleteState(id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, "State not found")
	}
	context.JSON(http.StatusNoContent, nil)
}
