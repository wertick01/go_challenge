package main

import (
	"fmt"
)

type Figure interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	length float64
	width float64
	area float64
	typ string
}

type Circle struct {
	radius float64
	area float64
	typ string
}

func Create_Rectangle(length float64, width float64) *Rectangle {
	return &Rectangle{area:length * width, typ: "Rectangle"}
}

func Create_Circle(radius float64) *Circle {
	return &Circle{area:float64(radius) * 3.14 * 2, typ: "Circle"}
}

func (r *Rectangle) Area() float64 {
	return r.area
}

func (c *Circle) Area() float64 {
	return c.area //!
}

func (r *Rectangle) Type() string {
	return r.typ
}

func (c *Circle) Type() string {
	return c.typ
}

func Result(f Figure) {
	fmt.Println("The figure type is", f.Type(), "and Area is:", f.Area())
}

func main() {
	rec := Create_Circle(14)
	Result(rec)

	recc := Create_Rectangle(12, 15)
	Result(recc)
}