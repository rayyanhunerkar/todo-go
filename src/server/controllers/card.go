package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/todo-go/src/database/models"
	"github.com/rayyanhunerkar/todo-go/src/database/repositories"
)

type CardController struct {
	service models.CardRepository
}

func InitCardController(cardRepo *repositories.CardRepo) *CardController {
	return &CardController{
		service: cardRepo,
	}
}

func (h *CardController) CreateCard(context *gin.Context) {
	var request models.CardRequest

	if err := context.BindJSON(&request); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.service.CreateCard(request, context.GetString("currentUser"), "id")
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusCreated, response)
}

func (h *CardController) GetCards(context *gin.Context) {

	response, err := h.service.GetCards(context.GetString("currentUser"))
	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, response)
}

func (h *CardController) GetCard(context *gin.Context) {
	id := context.Param("id")
	response, err := h.service.GetCardByID(id, context.GetString("currentUser"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
	}

	context.JSON(http.StatusOK, response)
}

func (h *CardController) DeleteCard(context *gin.Context) {
	id := context.Param("id")
	err := h.service.DeleteCard(id, context.GetString("currentUser"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
	}
	context.JSON(http.StatusNoContent, nil)
}
