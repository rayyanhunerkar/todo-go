package models

import (
	"time"

	"github.com/google/uuid"
)

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "public.users"
}

type User struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username  string    `gorm:"column:username;size:256;not null;unique" json:"username"`
	Password  string    `gorm:"column:password;size:256;not null;" json:"password"`
	FirstName string    `gorm:"column:first_name;size:64;not null;" json:"first_name"`
	LastName  string    `gorm:"column:last_name;size:64;not null;" json:"last_name"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime:true;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoUpdateTime:true;" json:"updated_at"`
}

type RegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRepository interface {
	CreateUser(user RegisterRequest) (*User, error)
}
