package main

import (
	"log/slog"

	"github.com/Futturi/internal"
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
	red := pkg.InitRedis(rcfg)
	repo := repository.NewRepository(red)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

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
