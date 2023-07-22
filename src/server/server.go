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
	RegisterAuthRoutes(router, repos.UserRepo, conf)
	RegisterStateRoutes(router, repos.StateRepo, conf)
	RegisterCardRoutes(router, repos.CardRepo, conf)
	router.Run()
}

func initSwagger(router *gin.Engine) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
