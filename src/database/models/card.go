package models

import (
	"time"

	"github.com/google/uuid"
)

func (Card) TableName() string {
	return "public.cards"
}

type Card struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string    `gorm:"column:title;not null;" json:"title"`
	Description string    `gorm:"column:description;not null;" json:"description"`
	Deadline    time.Time `gorm:"column:deadline;not null;" json:"deadline"`
	State       State     `gorm:"column:state_id;foreignKey:State;references:ID;" json:"state_id"`
	User        User      `gorm:"column:user_id;foreignKey:User;references:ID;" json:"user_id"`
	AssignedTo  User      `gorm:"column:assigned_to;foreignKey:User;references:ID;" json:"assigned_to"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime:true;" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoUpdateTime:true;" json:"updated_at"`
}

type CardRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	State       string    `json:"state_id" binding:"required"`
	AssignedTo  string    `json:"assigned_to" binding:"-"`
}

type UpdateCardRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	State       string    `json:"state_id" binding:"required"`
	AssignedTo  string    `json:"assigned_to" binding:"-"`
}

type CardRepository interface {
	CreateCard(request CardRequest, uid string) (*Response, error)
	GetCards(uid string) (*Response, error)
	GetCardByID(id string, uid string) (*Response, error)
	DeleteCard(id string, uid string) error
	UpdateCard(request UpdateCardRequest, id string, uid string) (*Response, error)
}
