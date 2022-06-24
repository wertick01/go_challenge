package checker

import (
	"context"
	"fmt"
	"sync"
	"time"
	"github.com/go-co-op/gocron"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Ping() string
	GetID() string
	Measurable
	Health(context.Context) bool
}

type Checker struct {
	slice []Checkable
	snc sync.Mutex
}

type CheckProblem struct {
	ProblemId []string
	ProblemHealth []string
}

func (check *Checker) Add(cbl Checkable) {
	check.snc.Lock()
	check.slice = append(check.slice, cbl)
	check.snc.Unlock()
}

func (check *Checker) Helper(chp *CheckProblem) {
	for _, value := range check.slice {
		chp.ProblemHealth = append(chp.ProblemHealth, value.GetID())
	}
}

func (check *Checker) Check(cont context.Context) *CheckProblem {

	var chp *CheckProblem
	check.snc.Lock()
	defer check.snc.Unlock()

	for _, value := range check.slice {
		if !value.Health(cont) {
			fmt.Println("--> ID", value.GetID(),  " doesn't working successfully")
			chp.ProblemId = append(chp.ProblemId, value.GetID())
		}
		continue
	}

	CheckPrinter(chp)
	return chp
}

func CheckPrinter(chp *CheckProblem) {
	for _, value := range(chp.ProblemId) {
		fmt.Printf(value)
	}
}

func Constructor(i, j int) *Checker {
	return &Checker{slice: make([]Checkable, i, j)}
}

func (check *Checker) Run(val int, cont context.Context, chp *CheckProblem) {
	//ticker := time.NewTicker(time.Duration(val) * time.Second)

	defer time.Sleep(7 * time.Second)

	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(func(){
		fmt.Printf("\n-->Process is running\n")
		check.Check(cont)
		go check.run(cont, chp)
	})

	for {
		select {
		case <- cont.Done():
			return
		}
	}
}

func (check *Checker) run(cont context.Context, chp *CheckProblem) {
	check.snc.Lock()
	defer check.snc.Unlock()

	for _, value := range check.slice {
		go func () {
			if !value.Health(cont) {
				check.Helper(chp)
			}
		} ()
	}
}

func (check *Checker) Stop(cancel context.CancelFunc) {
	cancel()
}

