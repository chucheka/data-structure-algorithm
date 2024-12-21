package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println(adder(1, 2, 3, 4))

	age := 20

	floatAge := float64(age)

	fmt.Println(floatAge)

	var i interface{} = 80

	//val, ok := i.(int) // asserting that the interface value is of type int
	//
	//fmt.Println(val, ok)

	if val, ok := i.(int); ok {
		fmt.Println(val, ok)
	}

	switch os := runtime.GOOS; os {
	case "mac":
		fmt.Println("It is a MAC")
	case "windows":
		fmt.Println("It is a windows")
	default:
		fmt.Println("It is not a mac or windows")
	}
}

func adder(nums ...int) int {

	var sum int = 0

	for _, num := range nums {
		sum += num
	}

	return sum
}
