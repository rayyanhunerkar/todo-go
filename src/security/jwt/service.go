package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rayyanhunerkar/todo-go/src/database/models"
	"github.com/spf13/viper"
)

type JWTService interface {
	GenerateToken(user models.User) (string, error)
	ValidateToken(c *gin.Context) (*jwt.Token, error)
	ExtractToken(c *gin.Context) string
}

type JWTConf struct {
	conf *viper.Viper
}

func InitJWTConf(conf *viper.Viper) *JWTConf {
	return &JWTConf{
		conf: conf,
	}
}

func (conf *JWTConf) GenerateToken(user models.User) (string, error) {
	secret := conf.conf.GetString("jwt.secret")
	claims := jwt.MapClaims{}
	claims["username"] = user.Username
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(60)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (conf *JWTConf) ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (conf *JWTConf) ValidateToken(c *gin.Context) (*jwt.Token, error) {
	secret := conf.conf.GetString("jwt.secret")
	tokenString := conf.ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
