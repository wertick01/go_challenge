package main

import (
	"context"
	"fmt"
	"time"

	"./internals/app"
	"./internals/checker"
	"./internals/config"
)

func main() {

	cst := Constructor(0, 1)
	cont, cancel := context.WithCancel(context.Background())
	cnfg := config.LoadAndStoreConfig()

	/*
	var client string
	var tm, count int
	fmt.Scanln("Введите количество клиентов: ",&count)
	for i := 0; i < count; i ++ {
		fmt.Scanf("%s\t", &client)
		fmt.Scanf("%d\n", &tm)
		cst.Add(Create_Client(client, time.Duration(tm) * time.Second))
	}	
	*/
	gmc_1 := CreateClient("pasha", 10 * time.Second)
	gmc_2 := CreateClient("", 2 * time.Second)
	gmc_3 := CreateClient("ne_pasha", 6 * time.Second)

	time.Sleep(2 * time.Second)
	
	cst.Add(gmc_1)
	cst.Add(gmc_2)
	cst.Add(gmc_3)

	cst.Check(cont)

	cst.Run(5, cont)

	newserv := app.NewServer(config, cont)
}