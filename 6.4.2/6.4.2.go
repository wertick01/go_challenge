package main

import (
	"flag"
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"os"
)

/*
Закрепим навык получения JSON-ответа в теле HTTP-запроса и его десериализацию
во внутреннюю структуру программы на Go.

Предположим, что где-то есть сервер, который при обращении по пути /health возвращает
следующую JSON-конструкцию, описывающую состояние работы какой-то системы

В этом задании нам необходимо написать программу, которая примет адрес сервера
по переданному в программу флагу --url, с помощью веб-клиента обратится по этому адресу, и
в случае успешного получения ответа этой структуры выведет на экране значения полей status,
service_id и checks.ping_mysql.status  в виде строки “Overall status is %status%, with service_id
%значение поля service_id% mysql component is  %начение поля checks.ping_mysql.status%”. А если по
какой-либо причине получить ответ от сервера не получилось, то вывести на экран "No data"
*/

type Responser struct {
	Status		string `json:"status"`
	ServiceID	string `json:"service_id"`
	Checks 		Checks `json:"checks"`
}

type Checks struct {
	PingMySQL 	PingMySQL `json:"ping_mysql"`
}

type PingMySQL struct {
	ComponentID 	string `json:"component_id"`
	ComponentType	string `json:"component_type"`
	Status			string `json:"status"`
}

func checker(err error) bool {
	if err != nil {
		fmt.Println("Thats panic !!")
		panic(err)
	}
	return true
}

func ResponserFunc(url string) *Responser {
	response, err_1 := http.Get(url)

	if err_1 != nil {
		fmt.Println("No data")
		checker(err_1)
	}

	result, err_2 := io.ReadAll(response.Body)

	checker(err_2)

	var resp *Responser

	json.Unmarshal(result, &resp)

	return resp
}

func ResponsePrinter(url string) {
	defer os.Exit(0)
	resp := ResponserFunc(url)

	fmt.Printf(
		"Overall status is %s, with service_id %s mysql component is  %s.", 
		resp.Status, resp.ServiceID, resp.Checks.PingMySQL.Status,
	)
}

func main() {
	addr := flag.String("url", "/", "url")
	flag.Parse()
	ResponsePrinter(*addr)
}
