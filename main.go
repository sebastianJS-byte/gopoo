package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Print(y,"\n")

	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println(myValue)
	}

	m := make(map[string]int)

	m["Key"] = 6
	
}
