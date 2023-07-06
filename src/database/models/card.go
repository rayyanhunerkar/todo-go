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
	State       State     `gorm:"column:state_id;foreignKey:StateRefer;" json:"state_id"`
	UserID      User      `gorm:"column:user_id;foreignKey:UserRefer;" json:"user_id"`
	AssignedTo  User      `gorm:"column:assigned_to;foreignKey:AssignRefer;" json:"assigned_to"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime:true;" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoUpdateTime:true;" json:"updated_at"`
}
