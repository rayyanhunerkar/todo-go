package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo  *UserRepo
	StateRepo *StateRepo
	CardRepo  *CardRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := InitUserRepo(db)
	stateRepo := InitStateRepo(db)
	cardRepo := InitCardRepo(db)
	return &Repositories{
		UserRepo:  userRepo,
		StateRepo: stateRepo,
		CardRepo:  cardRepo,
	}
}
