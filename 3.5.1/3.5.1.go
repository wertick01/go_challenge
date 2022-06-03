package main

import (
	"fmt"
	"math"
)

type Figure interface {
	Area(ar int) int
	Type(st string) string
}

type Rectangle struct {
	length int
	width int
	typ string
}

type Circle struct {
	radius int
	typ string
}

func (r *Rectangle) Area(ar int) int {
	//fmt.Println(r.length * r.width)
	return r.length * r.width
}

func (c *Circle) Area(ar int) int {
	return c.radius * int(math.Pow(3.14, 2)) //!
}

func (r *Rectangle) Type(st string) string {
	//fmt.Println(r.typ)
	return r.typ
}

func (c *Circle) Type(st string) string {
	return c.typ
}

func Result(f Figure) {
	fmt.Println(f.Area)
}

func main() {
	rec := &Rectangle{length: 14, width: 6, typ: "CIRCLE"}
	Result(rec)
}