package repositories

import (
	"github.com/rayyanhunerkar/todo-go/src/database/models"
	"gorm.io/gorm"
)

type StateRepo struct {
	db *gorm.DB
}

func InitStateRepo(db *gorm.DB) *StateRepo {
	return &StateRepo{
		db: db,
	}
}

func (repo *StateRepo) CreateState(request models.StateCreateRequest) (*models.State, error) {
	var err error
	var s models.State

	s.Name = request.Name
	s.Description = request.Description
	err = repo.db.Create(&s).Error

	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (repo *StateRepo) GetStates() (*[]models.State, error) {
	var states []models.State

	if result := repo.db.Find(&states); result.Error != nil {
		return nil, result.Error
	}
	return &states, nil
}

func (repo *StateRepo) GetStateByID(id string) (*models.State, error) {
	var s models.State

	if result := repo.db.Where("id = ?", id).First(&s); result.Error != nil {
		return nil, result.Error
	}
	return &s, nil
}
