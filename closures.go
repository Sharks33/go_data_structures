package main

import "fmt"

func main() {
	num1 := 1

	doubleNum := func() int {
		num1 *= 2
		return num1
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
}
