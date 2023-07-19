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

func (repo *StateRepo) CreateState(request models.StateCreateRequest) (*models.Response, error) {
	var err error
	var s models.State
	var response models.Response

	s.Name = request.Name
	s.Description = request.Description
	err = repo.db.Create(&s).Error
	if err != nil {
		return nil, err
	}

	response.Data = &s
	response.Message = "State created successfully"

	return &response, nil
}

func (repo *StateRepo) GetStates() (*models.Response, error) {
	var states []models.State
	var response models.Response

	if result := repo.db.Find(&states); result.Error != nil {
		return nil, result.Error
	}

	response.Data = &states
	response.Message = "Retrieved States successfully"
	return &response, nil
}

func (repo *StateRepo) GetStateByID(id string) (*models.Response, error) {
	var s models.State
	var response models.Response
	if result := repo.db.Where("id = ?", id).First(&s); result.Error != nil {
		return nil, result.Error
	}
	response.Data = &s
	response.Message = "State recieved successfully"
	return &response, nil
}

func (repo *StateRepo) UpdateState(request models.StateUpdateRequest, id string) (*models.Response, error) {
	var s models.State
	var response models.Response

	if result := repo.db.Where("id = ?", id).First(&s); result.Error != nil {
		return nil, result.Error
	}
	s.Name = request.Name
	s.Description = request.Description

	repo.db.Save(&s)
	response.Data = &s
	response.Message = "State updated successfully"
	return &response, nil
}

func (repo *StateRepo) DeleteState(id string) error {
	var s models.State

	if result := repo.db.Where("id = ?", id).Delete(&s); result.Error != nil {
		return result.Error
	}
	return nil
}
