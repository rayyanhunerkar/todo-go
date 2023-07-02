package repositories

import (
	"github.com/rayyanhunerkar/todo-go/src/database/models"
	"github.com/rayyanhunerkar/todo-go/src/utils"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func InitUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
func (repo *UserRepo) CreateUser(user models.RegisterRequest) (*models.User, error) {

	var err error
	var hashedPassword []byte
	var u models.User

	hashedPassword, err = utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Username = user.Username
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Password = string(hashedPassword)

	err = repo.db.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, err
}
