package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "съешь ещё этих мягких французских булок, да выпей чаю"
	/*
	res := make(map[string]int)

	for _, value := range str {
		res[value] = strings.Count(value)
	}
	fmt.Println(res)
	*/
	for _, value := range str {
		fmt.Println(value)
	}

}
