package service

import (
	"userBar/back/models"
	"userBar/back/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(name, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Client interface {
	GetById(userId int) (models.Me, error)
	GetList(userId int)([]models.Drinks, error)
	UpdateUser(userId, drinkId int) error
}

type Barman interface {
	GetList(barId int)([]models.Drinks, error)
	CreateDrink(barId int, name string, promille, cost float64)(int, error)
}

type Promiller interface {
	PromilleDec() string
}
type Service struct {
	Authorization
	Client
	Barman
	Promiller
}

func NewService(rep *repository.Repository) *Service  {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		Client: NewClientService(rep.Client),
		Barman: NewBarmanService(rep.Barman),
		Promiller: NewPromillerService(rep.Promiller),
	}
}