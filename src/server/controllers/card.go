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

	response, err := h.service.CreateCard(request, context.GetString("currentUser"))
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusCreated, response)
}
