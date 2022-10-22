package phonebooks

import "fmt"

func PhoneBooks() {
	phoneBooks := map[int]string{
		912334456:  "Shedrack",
		734569345:  "Guy pierce",
		237890345:  "Pierce brosnan",
		8176486028: "Jack sparrow",
	}

	for k, v := range phoneBooks {
		fmt.Println(k, "-", v)
	}

	fmt.Println(phoneBooks[734569345])
	phoneBooks[734569345] = "Jules"
	fmt.Println(phoneBooks)
}
