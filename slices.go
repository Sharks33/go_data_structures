package main

import "fmt"

func main() {

	// Slices are like arrays but you leave out the size
	numSlice := []int{5, 4, 3, 2, 1}

	// You can create a slice by defining the first index value to
	// take through the last
	numSlice2 := numSlice[3:5] // numSlice3 == [2,1]
	fmt.Println("numSlice2[0] =", numSlice2[0])

	// If you don't supply the first index it defaults to 0
	// If you don't supply the last index it defaults to max
	fmt.Println("numSlice[:2] =", numSlice[:2])
	fmt.Println("numSlice[2:] =", numSlice[2:])

	// You can also create an empty slice and define the data type,
	// length (receive value of 0), capacity (max size)
	numSlice3 := make([]int, 5, 10)

	// You can copy a slice to another
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3[0])

	// Append values to the end of a slice
	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3[6])
}
