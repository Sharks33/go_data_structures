package main

import "fmt"

func main() {

	legalDrinkingAge := 21
	myAge := 35

	if myAge > legalDrinkingAge {
		fmt.Println("I can have a craft beer")
	} else if myAge == legalDrinkingAge {
		fmt.Println("I can have a craft beer, just barely")
	} else {
		fmt.Println("I can't have a craft beer")
	}
}
