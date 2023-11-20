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
	cardController  *controllers.CardController
}

func RegisterPublicAuthRoutes(router *gin.Engine, userRepo *repositories.UserRepo, conf *viper.Viper) {

	h := &Controllers{
		userController: controllers.InitUserController(userRepo, conf),
	}

	routes := router.Group("/auth")
	routes.Use(middlewares.CORSMiddleware())
	routes.POST("/signup", h.userController.Register)
	routes.POST("/login", h.userController.Login)
}

func RegisterUserRoutes(router *gin.Engine, userRepo *repositories.UserRepo, conf *viper.Viper) {
	h := &Controllers{
		userController: controllers.InitUserController(userRepo, conf),
	}
	routes := router.Group("/user")
	routes.Use(middlewares.CORSMiddleware())
	routes.Use(middlewares.AuthJWTMiddleware(conf))
	routes.GET("/me", h.userController.Me)

}

func RegisterStateRoutes(router *gin.Engine, stateRepo *repositories.StateRepo, conf *viper.Viper) {

	h := &Controllers{
		stateController: controllers.InitStateController(stateRepo),
	}

	routes := router.Group("/states")
	routes.Use(middlewares.CORSMiddleware())
	routes.Use(middlewares.AuthJWTMiddleware(conf))
	routes.POST("", h.stateController.CreateState)
	routes.GET("", h.stateController.GetStates)
	routes.GET("/:id", h.stateController.GetStateByID)
	routes.PATCH("/:id", h.stateController.UpdateState)
	routes.DELETE("/:id", h.stateController.DeleteState)
}

func RegisterCardRoutes(router *gin.Engine, cardRepo *repositories.CardRepo, conf *viper.Viper) {
	h := &Controllers{
		cardController: controllers.InitCardController(cardRepo),
	}

	routes := router.Group("/cards")
	routes.Use(middlewares.CORSMiddleware())
	routes.Use(middlewares.AuthJWTMiddleware(conf))
	routes.POST("", h.cardController.CreateCard)
	routes.GET("", h.cardController.GetCards)
	routes.GET("/:id", h.cardController.GetCard)
}
