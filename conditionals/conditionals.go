package conditionals

import "fmt"

func Conditionals() {
	age := 45

	// fmt.Println(age >= 50)
	// fmt.Println(age <= 50)
	// fmt.Println(age == 40)
	// fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	socials := []string{"twitter", "snapchat", "instagram", "facebook", "tiktok", "whatsapp"}

	for index, social := range socials {
		if index == 1 {
			fmt.Println("Continue at position", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at position", index)
			break
		}
		fmt.Printf("the value at position %v is %v\n", index, social)
	}
}
