package repository

import (
	"github.com/jmoiron/sqlx"
	"userBar/back/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(name, password string) (models.User, error)
}

type Client interface {
	GetById(id int) (models.Me, error)
	GetList(userId int)([]models.Drinks, error)
	UpdateUser(userId, drinkId int) error
}

type Barman interface {
	GetList(barId int)([]models.Drinks, error)
	CreateDrink(barId int, name string, promille, cost float64)(int, error)

}
type Promiller interface {
	PromillerDec() string
}
type Repository struct {
	Authorization
	Client
	Barman
	Promiller
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: 	NewAuthPostgres(db),
		Client:			NewClient(db),
		Barman: 		NewBarman(db),
		Promiller: 		NewPromiller(db),
	}
}
