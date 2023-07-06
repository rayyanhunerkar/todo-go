package main

import (
	"github.com/rayyanhunerkar/todo-go/src/config"
	"github.com/rayyanhunerkar/todo-go/src/database"
	"github.com/rayyanhunerkar/todo-go/src/server"
)

func main() {
	conf := config.InitConfig()
	db := database.InitDatabase(conf)
	server.InitServer(db, conf)
}
