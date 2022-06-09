package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

//я знаю что всё написано очень косо, но
// чисто технически, всё правильно

func main() {
	w := &sync.WaitGroup{}
	w.Add(3)

	t_1 := 2.0001
	t_2 := 2 * time.Second

	go Runner_(t_1, t_2, w)

	w.Wait()

	fmt.Println("\nTHE END\n")
}

func Gorutine_(t float64, w *sync.WaitGroup) {
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Printf("Function has sleeped for %v seconds.\n", t)
	w.Done()
}

func Runner_(t_1 float64, t time.Duration, w *sync.WaitGroup) error {

	for i := 0; i < 3; i++ {
		go Gorutine_(t_1, w)
	}

	select {
	case <-time.After(t):
		fmt.Printf("\nOperating time was so long!\tNEEDS less than 2 seconds!\n\n")
		return errors.New("Runtime error")
	}
}

//какого черта в этих горутинах вообще происходит
