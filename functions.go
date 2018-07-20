package main

import "fmt"

func main() {
	numLst := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum:", addArr(numLst))

	// Get 2 values from a function
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)

	// Send an undefined number of values to a function (Variadic Function)
	fmt.Println(subtractThem(1, 2, 3, 4, 5))
}

func addArr(arrVal []float64) float64 {
	sum := 0.0

	for _, val := range arrVal {
		sum += val
	}

	return sum
}

// Go functions can return multiple values
func next2Values(number int) (int, int) {
	return number + 1, number + 2
}

// You can receive an undefined number of values with args ...int
func subtractThem(args ...int) int {
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}
