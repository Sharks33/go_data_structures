package main

import "fmt"

func main() {

	yourAge := 16

	switch yourAge {
	case 16:
		fmt.Println("You are old enough to drive")
	case 18:
		fmt.Println("You are old enough to vote")
	case 21:
		fmt.Println("You are old enough to drink")
	default:
		fmt.Println("Just how old are you really?")
	}
}
