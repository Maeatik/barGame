package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"userBar/back/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	var count int
	var passwords []string
	roleQuery := `SELECT COUNT(*) FROM users`
	rowQuery := a.db.QueryRow(roleQuery)

	if err := rowQuery.Scan(&count); err != nil {
		return 0, err
	}
	isUniq := true
	checkQuery := `SELECT password FROM users`
	if err := a.db.Select(&passwords, checkQuery); err != nil {
		return 0, err
	}

	for _, password := range passwords {
		if password == user.Password {
			isUniq = false
			fmt.Println(isUniq)
		}
	}

	query := `INSERT INTO users (name, password, money, promille, status, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	if isUniq {
		fmt.Println(123456789)
		if count == 0 {
			fmt.Println(1)
			fmt.Println(user.Name + " " + user.Password)
			row := a.db.QueryRow(query, user.Name, user.Password, 0, 0.0, "alive", false)
			if err := row.Scan(&id); err != nil {
				return 0, err
			}
			fmt.Println(2)
			return id, nil
		} else {
			row := a.db.QueryRow(query, user.Name, user.Password, 1000, 0.0, "alive", true)
			if err := row.Scan(&id); err != nil {
				return 0, err
			}
			return id, nil
		}
	} else {
		return 0, errors.New("password is busy")
	}

}

func (a *AuthPostgres) GetUser(name, password string) (models.User, error) {
	var user models.User
	fmt.Println(name + " " + password)
	query := fmt.Sprintf("SELECT id FROM users WHERE name=$1 AND password=$2")
	err := a.db.Get(&user, query, name, password)
	fmt.Println(user)
	return user, err
}
