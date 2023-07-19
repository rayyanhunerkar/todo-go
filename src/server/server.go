package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/todo-go/src/database/repositories"
	"github.com/rayyanhunerkar/todo-go/src/middlewares"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func InitServer(db *gorm.DB, conf *viper.Viper) {
	repos := repositories.InitRepositories(db)
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	RegisterAuthRoutes(router, repos.UserRepo, conf)
	RegisterStateRoutes(router, repos.StateRepo, conf)
	RegisterCardRoutes(router, repos.CardRepo, conf)
	router.Run()
}
