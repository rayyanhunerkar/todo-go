package main

import (
	"github.com/rayyanhunerkar/todo-go/src/config"
	"github.com/rayyanhunerkar/todo-go/src/database"
	"github.com/rayyanhunerkar/todo-go/src/server"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{}) // NEW
}

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	conf := config.InitConfig()
	db := database.InitDatabase(conf)
	server.Serve(db, conf)
}
