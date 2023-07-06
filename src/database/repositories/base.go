package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo  *UserRepo
	StateRepo *StateRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := InitUserRepo(db)
	stateRepo := InitStateRepo(db)
	return &Repositories{
		UserRepo:  userRepo,
		StateRepo: stateRepo,
	}
}
