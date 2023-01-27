package main

import (
	"Concurrency/examples"
	"fmt"
)

func main() {
	var input int
	fmt.Println("Enter the example number(1-8) you want to test:")
	fmt.Scanln(&input)

	switch input {
	case 1:
		examples.ExecuteOne()
	case 2:
		examples.ExecuteTwo()
	case 3:
		examples.ExecuteThree()
	default:
		fmt.Println("Wrong input. See you next time")
	}
}
