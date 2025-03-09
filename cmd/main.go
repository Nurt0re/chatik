package main

import (
	"log"
	"os"

	"github.com/Nurt0re/chatik"
	"github.com/Nurt0re/chatik/pkg/handler"
	"github.com/Nurt0re/chatik/pkg/repository"
	"github.com/Nurt0re/chatik/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	_"golang.org/x/oauth2/google"
)

type App struct{
	config *oauth2.Config
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading environmental variables %s", err.Error())
	}




	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error occured while connecting to db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(chatik.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while starting the server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
