package test

import "fmt"

func MyTest() {
	fmt.Println("What's your date of birth?.")
	var birthYear int
	fmt.Scan(&birthYear)

	fmt.Println("Enter the current year?.")
	var currentYear int
	fmt.Scan(&currentYear)

	getBirth(birthYear, currentYear)
}

func getBirth(birthYear int, currentYear int) (int, int) {
	myAge := currentYear - birthYear
	fmt.Printf("You are %v years old\n", myAge)
	return birthYear, currentYear
}
