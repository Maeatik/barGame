package service

import (
	"userBar/back/models"
	"userBar/back/pkg/repository"
)

type ClientService struct {
	rep repository.Client
}


func NewClientService(rep repository.Client) *ClientService  {
	return &ClientService{rep: rep}
}

func (c *ClientService) GetById(userId int) (models.Me, error) {
	return c.rep.GetById(userId)
}
func (c *ClientService) GetList(userId int) ([]models.Drinks, error) {
	return c.rep.GetList(userId)
}

func (c *ClientService) UpdateUser(userId, drinkId int) error{
	return  c.rep.UpdateUser(userId, drinkId)
}