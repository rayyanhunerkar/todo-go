package main

import (
	"github.com/rayyanhunerkar/todo-go/src/config"
	"github.com/rayyanhunerkar/todo-go/src/database"
	"github.com/rayyanhunerkar/todo-go/src/server"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	conf := config.InitConfig()
	db := database.InitDatabase(conf)
	server.Serve(db, conf)
}
