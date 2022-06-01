package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "съешь ещё этих мягких французских булок, да выпей чаю"
	res := make(map[string]int)

	for _, value := range str {
		res[string(value)] = strings.Count(str, string(value))
	}
	fmt.Println(res)

}
