package functions

import (
	"fmt"
	"strings"
)

func MultipleFunc() {
	firstName, secondName := getInitials("Lockheed Martin")
	fmt.Println(firstName, secondName)
}

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
