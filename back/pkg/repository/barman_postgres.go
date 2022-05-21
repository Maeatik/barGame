package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"userBar/back/models"
)

type BarmanPostgres struct {
	db *sqlx.DB
}

func NewBarman(db *sqlx.DB) *BarmanPostgres  {
	return &BarmanPostgres{db: db}
}

func (b *BarmanPostgres) GetList(barId int)([]models.Drinks, error)  {
	var list []models.Drinks
	var role models.CheckRole

	queryRole := `SELECT role from USERS WHERE id=$1`
	if err := b.db.Get(&role, queryRole, barId); err != nil{
		return nil, err
	}

	if !role.Role{
		query := `SELECT item,cost,promille FROM drinks`
		err := b.db.Select(&list, query)
		fmt.Println(err)
		return list, err
	} else {
		return nil, errors.New("access is denied")
	}

}

func (b *BarmanPostgres) CreateDrink(barId int, name string, promille, cost float64)(int, error) {
	var id int
	var role models.CheckRole

	queryRole := `SELECT role from USERS WHERE id=$1`
	if err := b.db.Get(&role, queryRole, barId); err != nil{
		return 0, err
	}
	if !role.Role {
		query := `INSERT INTO drinks (item, cost, promille) VALUES ($1, $2, $3) RETURNING id`
		row := b.db.QueryRow(query, name, cost, promille)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
		return id, nil
	} else {
		return 0, errors.New("access is denied")
	}
}