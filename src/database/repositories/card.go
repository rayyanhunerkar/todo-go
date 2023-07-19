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

	card.Title = request.Title
	card.Description = request.Description
	card.Deadline = request.Deadline
	card.User = *user

	data.Data = &card
	data.Message = "Card created successfully"

	return &data, nil
}

func (repo *CardRepo) GetCards(uid string) (*models.Response, error) {

	return nil, nil
}

func (repo *CardRepo) GetCardByID(id string, uid string) (*models.Response, error) {
	return nil, nil
}

func (repo *CardRepo) DeleteCard(id string) {}
