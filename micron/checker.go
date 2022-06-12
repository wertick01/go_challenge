package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Ping() string
	GetID() string
	Health(context.Context) bool
	Measurable
}

type Checker struct {
	slice []Checkable
	sync.Mutex
}

type Checker_problem struct {
	Problem string
}

func (check *Checker) Add(c Checkable) {

	check.Lock()
	check.slice = append(check.slice, c)
	check.Unlock()

	return
}

func (check *Checker) Helper() []string {
	str := []string{}

	for _, value := range check.slice {
		str = append(str, value.GetID())
	}

	return str
}

func (check *Checker) Check(cont context.Context) []string {
	var res string
	check.Lock()
	defer check.Unlock()
	mass := []string{}
	for _, value := range check.slice {
		if value.Health(cont) == true {
			continue
		} else {
			res = "--> ID" +value.GetID() + " doesn't working successfully\n"
			//fmt.Println("--> ID", value.GetID(),  " doesn't working successfully\n")
			mass = append(mass, res)
		}
		Check_printer(mass)
	}
	//Check_printer(mass)
	return mass
}

func Check_printer(mass []string) {
	for _, value := range(mass) {
		fmt.Printf(value)
	}
	return
}

func Problem_printer(chp *Checker_problem) {
	fmt.Println(chp.Problem)
}

func Constructor(i, j int) *Checker {
	return &Checker{slice: make([]Checkable, i, j)}
}

func (check *Checker) Stop(cont context.CancelFunc) *Checker_problem {
	defer cont()
	var chp *Checker_problem
	chp.Problem = "-->Work has been STOPPED"
	//fmt.Println("Func Stop is working successfully")
	return chp
}

func (check *Checker) run(cont context.Context) {
	check.Lock()

	for _, value := range check.slice {
		go value.Health(cont)
	}

	check.Unlock()
}

func Create_ticker(val int) *time.Ticker {
	return time.NewTicker(time.Duration(val) * time.Second)
}

func (check *Checker) Run(val int, cont context.Context) {
	ticker := Create_ticker(val)
	defer ticker.Stop()

	for {
		select {
		case <- cont.Done():
			fmt.Printf("-->Process is very long. Sorry :(\n")
			return
		case <- ticker.C:
			fmt.Printf("\n-->Process is running\n")
			go check.run(cont)
		}
	}
}

