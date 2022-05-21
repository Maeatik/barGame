package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"userBar/back/models"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClient(db *sqlx.DB) *ClientPostgres  {
	return &ClientPostgres{db: db}
}
func (c *ClientPostgres) GetById(userId int)(models.Me, error)  {
	var list models.Me
	var role models.CheckRole

	queryRole := `SELECT role from USERS WHERE id=$1`
	if err := c.db.Get(&role, queryRole, userId); err != nil{
		return list, err
	}
	if role.Role{
		query := `SELECT money, promille, status FROM users where id=$1`
		err := c.db.Get(&list, query, userId)
		fmt.Println(err)
		return list, err
	} else {
		return list, errors.New("access is denied")
	}

}

func (c *ClientPostgres) GetList(userId int)([]models.Drinks, error)  {
	var list []models.Drinks
	var id models.CheckStatus
	var role models.CheckRole

	queryRole := `SELECT role from USERS WHERE id=$1`
	if err := c.db.Get(&role, queryRole, userId); err != nil{
		return nil, err
	}
	if role.Role {
		queryId := `SELECT status FROM users where id=$1`
		if err := c.db.Get(&id, queryId, userId); err != nil {
			return nil, err
		}
		if id.Status == "alive" {
			query := `SELECT item,cost,promille FROM drinks`
			err := c.db.Select(&list, query)
			return list, err
		} else {
			return nil, errors.New("you are dead, im so sorry")
		}
	} else {
		return nil, errors.New("access is denied")
	}
}

func (c *ClientPostgres) UpdateUser(userId, drinkId int) error  {
	var client models.Me
	var drinkItem models.DrinksUpdateInput
	var role models.CheckRole

	queryRole := `SELECT role from USERS WHERE id=$1`
	if err := c.db.Get(&role, queryRole, userId); err != nil{
		return err
	}
	if role.Role {
		args := make([]interface{}, 0)

		queryDrink := `SELECT cost, promille FROM drinks WHERE id = $1`
		if err := c.db.Get(&drinkItem, queryDrink, drinkId); err != nil{
			return err
		}

		queryClient := `SELECT money, promille, status from users WHERE id = $1`
		if err := c.db.Get(&client, queryClient, userId); err != nil{
			return err
		}

		if client.Status == "alive"{
			if client.Money > drinkItem.Cost{
				client.Money = client.Money - drinkItem.Cost
				client.Promille = client.Promille + drinkItem.Promille

				if client.Promille >= 6 {
					client.Status = "dead"
				}
				queryPostClient := `UPDATE users SET (money, promille, status) = ($1, $2, $3) WHERE id = $4`
				args = append(args, client.Money, client.Promille, client.Status, userId)
				_, err := c.db.Exec(queryPostClient, args...)
				return err
			} else{
				return errors.New("not enough money")
			}
		} else {
			return errors.New("you are dead, im so sorry")
		}
	} else {
		return  errors.New("access is denied")
	}

}