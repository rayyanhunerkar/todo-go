package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/todo-go/src/database/repositories"
	"github.com/rayyanhunerkar/todo-go/src/middlewares"
	"github.com/rayyanhunerkar/todo-go/src/server/controllers"

	"github.com/spf13/viper"
)

type Controllers struct {
	userController  *controllers.UserController
	stateController *controllers.StateController
}

func RegisterAuthRoutes(router *gin.Engine, userRepo *repositories.UserRepo, conf *viper.Viper) {

	h := &Controllers{
		userController: controllers.InitUserController(userRepo, conf),
	}

	routes := router.Group("/auth")
	routes.POST("/signup", h.userController.Register)
	routes.POST("/login", h.userController.Login)

}

func RegisterStateRoutes(router *gin.Engine, stateRepo *repositories.StateRepo, conf *viper.Viper) {

	h := &Controllers{
		stateController: controllers.InitStateController(stateRepo),
	}
	routes := router.Group("/states")
	routes.Use(middlewares.AuthJWTMiddleware(conf))
	routes.GET("/", h.stateController.GetStates)
	routes.POST("/", h.stateController.CreateState)
	routes.GET("/:id", h.stateController.GetStateByID)
	routes.PATCH("/:id")
	routes.DELETE("/:id")
}

func RegisterCardRoutes(router *gin.Engine, repositories *repositories.Repositories) {}