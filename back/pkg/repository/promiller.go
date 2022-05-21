package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
	"userBar/back/models"
)

type PromillerPostgres struct {
	db *sqlx.DB
}



func NewPromiller(db *sqlx.DB) *PromillerPostgres  {
	return &PromillerPostgres{db: db}
}

func (p PromillerPostgres) PromillerDec() string {
	var client []models.PromilleDecrease
	args := make([]interface{}, 0)
	go func() {
		for true {
			//минута реального времени = час игрового времени
			time.Sleep(10 * time.Second)
			query := `SELECT id, promille, status, role FROM users`
			if err := p.db.Select(&client, query); err != nil {
				logrus.Println(err)
			}
			for _, user := range client {
				if user.Role && user.Status == "alive" {
					if user.Promille > 1 {
						user.Promille = user.Promille - 1
					} else if user.Promille <= 1 {
						user.Promille = 0
					}
					queryPromiller := `UPDATE users SET promille = $1 WHERE id = $2`
					args = append(args, user.Promille, user.Id)
					_, err := p.db.Exec(queryPromiller, args...)
					args = nil
					logrus.Println(err)
				}
			}
		}
	}()
	return "Прошел час"
}
