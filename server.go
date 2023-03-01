package main

import (
	"context"
	"github.com/gusleein/golog"
	"goTest/api"
	"goTest/configs"
	"goTest/db"
	"os"
	"os/signal"
	"syscall"
)

//api - рест апи
//db - коннектор к бд и модели
//heplers - вспомогательные ф-ии
//api.go - стартовый файл

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configs.Init()
	log.Init(true, log.Console)
	db.Init(ctx)

	go api.Start(ctx)
	gracefulShutDown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutDown, syscall.SIGINT, syscall.SIGTERM)

	<-gracefulShutDown
}
