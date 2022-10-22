package arrays

import "fmt"

func Arrays() {
	var cars [4]string
	cars[0] = "BMW"
	cars[1] = "Porsche"
	cars[2] = "Benz"
	cars[3] = "Toyota"

	for _, index := range cars {
		fmt.Printf("Cars: %q\n", index)
	}

	var ages [10]int = [10]int{20, 25, 45, 12, 15, 22, 23, 21, 29, 10}

	for _, age := range ages {
		fmt.Printf("User Ages: %v\n", age)
	}

	var users = [3]string{"Maya", "Jake", "Isabele"}

	for _, user := range users {
		fmt.Printf("Users in qoute: %q\n", user)
	}

	fmt.Printf("Users length: %v\n", len(users))

	cities := [4]string{"Mumbai", "Delhi", "Goa", "Locknow"}

	for _, citiesInIndia := range cities {
		fmt.Printf("Cities: %v\n", citiesInIndia)
	}
}
