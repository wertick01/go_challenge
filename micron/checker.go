package main

import "fmt"

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Ping() string
	GetID() string
	Health() bool
	Measurable
}

type Checker struct {
	slice []Checkable
}

func (check *Checker) Add(c Checkable) {
	check.slice = append(check.slice, c)
}

func (check Checker) Helper() string {
	str := ""

	for _, value := range check.slice {
		str += value.GetID()
	}

	return str
}

func (check *Checker) Check() {
	for _, value := range check.slice {
		if value.Health() {
			continue
		} else {
			fmt.Println(value.GetID() + " не работает")
		}
	}
}

func Constructor() *Checker {
	return &Checker{slice: make([]Checkable, 0, 1)}
}
