package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(mass_2 []int, val int) []int {
	mass_2 = append(mass_2, val)
	fmt.Printf("--> Value #%v has been appended to final massive\n", val)
	return mass_2
}

func main() {
	var pong, ping sync.Mutex
	a := [5]int{1, 3, 5, 7, 9}
	b := [5]int{2, 4, 6, 8, 10}
	c := []int{}

	pong.Lock()
	go func() {
		for _, val := range a {
			ping.Lock()
			c = foo(c, val)
			pong.Unlock()
		}
	}()

	go func() {
		for _, val := range b {
			pong.Lock()
			c = foo(c, val)
			ping.Unlock()
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("The result is \t-->%v<--\n", c)
}
