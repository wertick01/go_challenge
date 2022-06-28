package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Smth_function(cont context.Context, w *sync.WaitGroup, num int, i int) {
	defer w.Done()
	tm := time.NewTimer(time.Duration(num) * time.Second)
	select {
	case <-cont.Done():
		fmt.Printf("-->Process %v have been stoped :((\n", i+1)
	case <-tm.C:
		fmt.Printf("-->Process %v done :))\n", i+1)
	}
}

func Smth_runner(num int) {
	cont, cancel := context.WithCancel(context.Background())
	//defer cancel()
	w := new(sync.WaitGroup)
	defer w.Wait()
	defer cancel()
	w.Add(num)
	for i := 0; i < num; i++ {
		go Smth_function(cont, w, num, i)
	}
	time.Sleep(3 * time.Second)
}

func main() {
	Smth_runner(2)
}

/*
Напишите функцию, которая запускает несколько горутин, передавая в них контекст.
Потом через несколько секунд (например 5) функция отменяет контекст, и созданные горутины
должны отреагировать на отмену контекста: вывести на экран сообщение и завершиться.
*/
