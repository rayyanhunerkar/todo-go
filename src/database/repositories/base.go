package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo *UserRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := InitUserRepo(db)
	return &Repositories{
		UserRepo: userRepo,
	}
}
