package main

import (
	"log/slog"

	"github.com/Futturi/internal"
	"github.com/Futturi/internal/auth_service"
	"github.com/Futturi/internal/handler"
	"github.com/Futturi/internal/repository"
	"github.com/Futturi/internal/service"
	"github.com/Futturi/pkg"
	"github.com/spf13/viper"
)

func main() {
	err := InitConfig()
	if err != nil {
		slog.Error(err.Error())
	}
	rcfg := pkg.RedisConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DB:       0,
		Password: "",
	}

	pcfg := pkg.Config{
		Hostname: viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DB:       viper.GetString("db.namedb"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := pkg.InitPostges(pcfg)
	if err != nil {
		slog.Error(err.Error())
	}

	grpcclient, err := authservice.InitAuth(viper.GetString("grpc.host"), viper.GetString("grpc.port"))
	if err != nil {
		slog.Error(err.Error())
	}
	red := pkg.InitRedis(rcfg)
	repo := repository.NewRepository(red, db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service, grpcclient)
	serv := new(internal.Server)
	if err := serv.InitRoutes(viper.GetString("port"), handler.Init()); err != nil {
		slog.Error(err.Error())
	}

}
func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("internal/config")
	return viper.ReadInConfig()
}
