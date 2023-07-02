package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/todo-go/src/database/repositories"
)

type Controllers struct {
	userController *UserController
}

func RegisterAuthRoutes(router *gin.Engine, repositories *repositories.Repositories) {
	h := &Controllers{
		userController: InitController(repositories.UserRepo),
	}
	routes := router.Group("/users")
	routes.POST("/signup", h.userController.register)
}
