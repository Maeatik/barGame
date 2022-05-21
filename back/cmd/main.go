package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
	"userBar/back"
	"userBar/back/pkg/handler"
	"userBar/back/pkg/repository"
	"userBar/back/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil{
		logrus.Fatalf("error initializating configs: %s", err.Error())
	}
	if err := gotenv.Load(); err != nil{
		logrus.Fatalf("error initializating env file: %s", err.Error())
	}
	db, err := repository.OpenConnection(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil{
		logrus.Fatalf("error initializating configs: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	hour := services.Promiller.PromilleDec()
	fmt.Println(hour)

	server := new(back.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err!= nil{
		logrus.Fatalf("err while running http server")
	}
}

func initConfig() error  {
	viper.AddConfigPath("back\\config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
