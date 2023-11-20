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

func (repo *UserRepo) CreateUser(request models.RegisterRequest) (*models.Response, error) {

	var err error
	var hashedPassword []byte
	var u models.User
	var data models.RegisterResponse
	var response models.Response

	hashedPassword, err = utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	u.Username = request.Username
	u.FirstName = request.FirstName
	u.LastName = request.LastName
	u.Password = string(hashedPassword)

	err = repo.db.Create(&u).Error
	if err != nil {
		return nil, err
	}

	data.ID = u.ID
	data.FirstName = u.FirstName
	data.LastName = u.LastName
	data.Username = u.Username

	response.Data = data
	response.Message = "User Created successfully"

	return &response, nil
}

func (repo *UserRepo) Login(request models.LoginRequest, conf *viper.Viper) (*models.Response, models.ErrorResponse) {
	var data models.LoginResponse
	var response models.Response
	var errorResponse models.ErrorResponse

	jwtConf := jwt.InitJWTConf(conf)

	result, err := repo.GetUserByEmail(request.Username)
	if err != nil {
		errorResponse.Message = "Wrong Username/Password"
		return nil, errorResponse
	}
	err = utils.VerifyPassword(result.Password, request.Password)
	if err != nil {
		errorResponse.Message = "Wrong Username/Password"
		return nil, errorResponse
	}
	token, err := jwt.JWTService.GenerateToken(jwtConf, *result)
	if err != nil {
		panic(err)
	}

	data.AccessToken = token
	data.ID = result.ID
	data.Username = result.Username

	response.Data = data
	response.Message = "User Logged in successfully"
	return &response, errorResponse
}

func (repo *UserRepo) GetUserByID(id string) (*models.User, error) {
	var u models.User
	if result := repo.db.Where("id = ?", id).First(&u); result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}

func (repo *UserRepo) GetUserByEmail(username string) (*models.User, error) {
	var u models.User
	if result := repo.db.Where("username = ?", username).First(&u); result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}
