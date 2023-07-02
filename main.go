package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/todo-go/src/config"
	"github.com/rayyanhunerkar/todo-go/src/controller/auth"
	"github.com/rayyanhunerkar/todo-go/src/database"
	"github.com/rayyanhunerkar/todo-go/src/database/repositories"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.InitDatabase(conf)
	if err != nil {
		panic(err)
	}
	repos := repositories.InitRepositories(db)
	router := gin.Default()

	auth.RegisterAuthRoutes(router, repos)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port": 8080,
		})
	})
	router.Run()
}
