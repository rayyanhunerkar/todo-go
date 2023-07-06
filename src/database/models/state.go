package models

import (
	"time"

	"github.com/google/uuid"
)

func (State) TableName() string {
	return "public.states"
}

type State struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"column:name;not null;" json:"name"`
	Description string    `gorm:"column:description;not null;" json:"description"`
	CreatedOn   time.Time `gorm:"column:created_on;not null;autoCreateTime:true;" json:"created_at"`
	ModifiedOn  time.Time `gorm:"column:modified_on;not null;autoUpdateTime:true;" json:"updated_at"`
}

type StateCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type StateUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type StateRepository interface {
	CreateState(request StateCreateRequest) (*State, error)
	GetStates() (*[]State, error)
	GetStateByID(id string) (*State, error)
}