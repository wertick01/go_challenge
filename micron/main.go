package main

import (
	"fmt"
)

func main() {
	check := Constructor()


	gmc_1 := Create_Client("pasha")
	gmc_2 := Create_Client("toje_pasha")
	gmc_3 := Create_Client("ne_pasha")


	check.Add(gmc_1)
	check.Add(gmc_2)
	check.Add(gmc_3)

	fmt.Print(check)

	check.Check()
}