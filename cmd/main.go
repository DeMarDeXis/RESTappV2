package main

// 5 31
import (
	"os"

	gorestapiv2 "github.com/DeMarDeXis/RESTV1"
	"github.com/DeMarDeXis/RESTV1/pkg/handler"
	"github.com/DeMarDeXis/RESTV1/pkg/service"
	"github.com/DeMarDeXis/RESTV1/pkg/storage"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars: %s", err.Error())
	}

	db, err := storage.NewPostgresDB(storage.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error occured while connecting to db: %s", err.Error())
	}

	storage := storage.NewStorage(db)
	services := service.NewService(storage)
	handlers := handler.NewHandler(services)

	srv := new(gorestapiv2.Server)
	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
