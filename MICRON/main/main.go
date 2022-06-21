package main

import (
	app "./MICRON/internals/app"
	cnf "./MICRON/internals/config"
	"context"
)

func main() { //точка входа нашего сервера
	config := cnf.LoadAndStoreConfig() //грузим конфигурацию

	ctx, _ := context.WithCancel(context.Background()) // создаем контекст для работы контекстнозависимых частей системы

	server := app.NewServer(config, ctx) // создаем сервер

	param := "str"

	server.Serve(ctx, param) //запускаем сервер

	server.Shutdown()

}
