package main

import (
	"fmt";
	"sort"
)

func main() {

	var names = []string{"st", "a", "res", "c", "daily"}
    sort.Strings(names)
    fmt.Println("Sorted in alphabetical order", names)

}
