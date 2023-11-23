package repositories

import (
	"github.com/rayyanhunerkar/todo-go/src/database/models"
	"gorm.io/gorm"
)

type CardRepo struct {
	db *gorm.DB
}

func InitCardRepo(db *gorm.DB) *CardRepo {
	return &CardRepo{
		db: db,
	}
}

func (repo *CardRepo) CreateCard(request models.CardRequest, uid string) (*models.Response, error) {
	var card models.Card
	var data models.Response

	user, err := models.UserRepository.GetUserByID(InitUserRepo(repo.db), uid)
	if err != nil {
		return nil, err
	}

	state, err := models.StateRepository.GetState(InitStateRepo(repo.db), request.State)
	if err != nil {
		return nil, err
	}

	assignedUser, err := models.UserRepository.GetUserByID(InitUserRepo(repo.db), uid)
	if err != nil {
		return nil, err
	}

	card.Title = request.Title
	card.Description = request.Description
	card.Deadline = request.Deadline
	card.User = *user
	card.State = *state
	card.AssignedTo = *assignedUser

	data.Data = &card
	data.Message = "Card created successfully"

	return &data, nil
}

func (repo *CardRepo) GetCards(uid string) (*models.Response, error) {
	var cards []models.Card
	var data models.Response

	user, err := models.UserRepository.GetUserByID(InitUserRepo(repo.db), uid)
	if err != nil {
		return nil, err
	}

	if result := repo.db.Find(&cards).Where("user_id = ?", user.ID); result.Error != nil {
		return nil, err
	}

	data.Data = &cards
	data.Message = "Cards fetched successfully"
	return &data, nil

}

func (repo *CardRepo) GetCardByID(id string, uid string) (*models.Response, error) {
	var card models.Card
	var data models.Response

	user, err := models.UserRepository.GetUserByID(InitUserRepo(repo.db), uid)
	if err != nil {
		return nil, err
	}

	if result := repo.db.Find(&card).Where("user_id = ?", user.ID).Where("id = ?", id); result.Error != nil {
		return nil, err
	}

	data.Data = &card
	data.Message = "Card fetched successfully"
	return &data, nil

}

func (repo *CardRepo) DeleteCard(id string, uid string) error {
	var card models.Card

	user, err := models.UserRepository.GetUserByID(InitUserRepo(repo.db), uid)
	if err != nil {
		return err
	}

	if result := repo.db.Find(&card).Where("user_id = ?", user.ID).Where("id = ?", id).Delete(&card); result.Error != nil {
		return err
	}

	return nil
}

func (repo *CardRepo) UpdateCard(request models.UpdateCardRequest, id string, uid string) (*models.Response, error) {

	var card models.Card
	var data models.Response

	user, err := models.UserRepository.GetUserByID(InitUserRepo(repo.db), uid)
	if err != nil {
		return nil, err
	}

	state, err := models.StateRepository.GetState(InitStateRepo(repo.db), request.State)
	if err != nil {
		return nil, err
	}

	assignedUser, err := models.UserRepository.GetUserByID(InitUserRepo(repo.db), request.AssignedTo)
	if err != nil {
		return nil, err
	}

	if result := repo.db.Find(&card).Where("user_id = ?", user.ID).Where("id = ?", id).First(&card); result.Error != nil {
		return nil, err
	}

	card.Title = request.Title
	card.Description = request.Description
	card.State = *state
	card.AssignedTo = *assignedUser

	repo.db.Save(&card)

	data.Data = &card
	data.Message = "Task Updated Successfully"

	return &data, nil
}
