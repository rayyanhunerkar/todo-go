package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(conf *viper.Viper) (*gorm.DB, error) {

	dbName := conf.GetString("database.name")
	dbUser := conf.GetString("database.user")
	dbPassword := conf.GetString("database.password")
	dbHost := conf.GetString("database.host")
	dbPort := conf.GetString("database.port")

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dbURL))
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}