package repositories

import (
	"github.com/rayyanhunerkar/todo-go/src/database/models"
	"github.com/rayyanhunerkar/todo-go/src/security/jwt"
	"github.com/rayyanhunerkar/todo-go/src/utils"
	"github.com/spf13/viper"
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

	hashedPassword, err = utils.HashPassword(user.Password)
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

func (repo *UserRepo) Login(request models.LoginRequest, conf *viper.Viper) (string, error) {
	var u models.User
	jwtConf := jwt.InitJWTConf(conf)
	if result := repo.db.Where("username = ?", request.Username).First(&u); result.Error != nil {
		return "", result.Error
	}

	err := utils.VerifyPassword(u.Password, request.Password)
	if err != nil {
		return "", err
	}
	token, err := jwt.JWTService.GenerateToken(jwtConf, u)
	if err != nil {
		panic(err)
	}
	return token, nil
}
