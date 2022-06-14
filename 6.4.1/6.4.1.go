package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

/*
Для того, чтобы освоить основы пакетов net/http и os, напишите http сервер с
эндпоинтом POST /log,который принимает строку и помещает её в новую линию в
файле, путь к которому определен переменной env APP_LOGFILE_PATH.

Если путь не определен, то сервер должен использовать свой корень, где он запущен
с именем файла log.txt по умолчанию. После записи строки в лог сервер должен вернуть
код 200 и OK
*/

type Handler struct {
	writer http.ResponseWriter
	req    *http.Request
} // я хотел как-то применить эту структуру, шоб было красиво, но не получилось

func main() {
	os.Setenv("APP_LOGFILE_PATH", "log.txt")
	http.HandleFunc("/", FileWriter)
	http.ListenAndServe(":8080", nil)
}

func check(err error) {
	if err != nil {
		fmt.Println("Thats panic!!")
		panic(err)
	}
}

func FileWriter(wrt http.ResponseWriter, req *http.Request) {
	logpath := "../6.4.1/log.txt"

	if req.Method == "POST" {
		txt := "something\n"

		if !exister(logpath) {
			os.Create(logpath)
		}

		file, err := os.Open(logpath)
		check(err)

		writer := bufio.NewWriter(file)
		writer.WriteString(txt)
		writer.Flush()

		file.Close()

		go func() {
			wrt.WriteHeader(http.StatusOK)
			wrt.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			resp["message"] = "Status OK"
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			wrt.Write(jsonResp)
		}() // а тут я просто выпендриваюсь с горутиной
		time.Sleep(1 * time.Second)
	}
}

func exister(path string) bool {
	_, err := os.Stat(path)
	switch {
	case os.IsNotExist(err):
		fmt.Println("log file not exists... creating log.txt")
		return false
	case err != nil:
		fmt.Println(err)
		return false
	default:
		return true
	}
}
