package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const(
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "game_bar"
)

type Config struct {
	Host 		string
	Port 		string
	Username 	string
	Password 	string
	DBName 		string
}

type PostgresRepository struct {
	Db *sqlx.DB
}

func OpenConnection(cfg Config) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		logrus.Println(err)
	}
	//Проверка подключения
	err = db.Ping()
	if err != nil {
		logrus.Println(err)
	}
	//Возвращение указателя на новую БД
	return db, err
}

func (r *PostgresRepository) Close() {
	r.Db.Close()
}

func (r *PostgresRepository) CreateUser()  {

}