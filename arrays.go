package main

import "fmt"

func main() {

	// Example 1
	var myArr [3]float64

	myArr[0] = 17
	myArr[1] = 29
	myArr[2] = 233

	fmt.Println(myArr)
	fmt.Println(myArr[2])

	// Example 2
	myArr2 := [3]float64{1, 2, 3}

	fmt.Println(myArr2)
	fmt.Println(myArr2[2])

	// Loop through array and print index and value
	for i, val := range myArr2 {
		fmt.Println(i, val)
	}
}
