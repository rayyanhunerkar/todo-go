package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/todo-go/docs"
	"github.com/rayyanhunerkar/todo-go/src/database/repositories"
	"github.com/rayyanhunerkar/todo-go/src/middlewares"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Serve(db *gorm.DB, conf *viper.Viper) {
	repos := repositories.InitRepositories(db)
	router := gin.Default()
	initSwagger(router)
	router.Use(middlewares.CORSMiddleware())
	RegisterPublicAuthRoutes(router, repos.UserRepo, conf)
	RegisterStateRoutes(router, repos.StateRepo, conf)
	RegisterCardRoutes(router, repos.CardRepo, conf)
	RegisterUserRoutes(router, repos.UserRepo, conf)
	router.Run("0.0.0.0:8000")
}

func initSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Todo"
	docs.SwaggerInfo.Version = "v1"
	url := ginSwagger.URL("http://0.0.0.0:8000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
