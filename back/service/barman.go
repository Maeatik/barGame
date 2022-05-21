package service

import (
	"userBar/back/models"
	"userBar/back/pkg/repository"
)

type BarmanService struct {
	rep repository.Barman
}

func NewBarmanService(rep repository.Barman) *BarmanService  {
	return &BarmanService{rep: rep}
}

func (b *BarmanService) GetList(barId int)([]models.Drinks, error){
	return b.rep.GetList(barId)
}
func (b *BarmanService) CreateDrink(barId int, name string, promille, cost float64)(int, error){
	return b.rep.CreateDrink(barId, name, promille, cost)
}
