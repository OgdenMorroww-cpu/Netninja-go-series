package functions

import (
	"fmt"
	"math"
)

func Mainfunc() {
	sayGreetings("Mason")
	sayGreetings("Axwell")
	sayBye("Martha")
	sayBye("Lois")
	cycleNames([]string{"Mason", "Stefan", "Damon", "Bonnie", "Matt", "Jeremy", "Josh"}, sayGreetings)
	cycleNames([]string{"Wade", "Barret", "Scott", "Adkins", "Jean claude"}, sayBye)
	area := cycleArea(120)
	fmt.Printf("The radius is: %0.2f\n", area)
}

func sayGreetings(greet string) {
	fmt.Printf("Good morning %v\n", greet)
}

func sayBye(bye string) {
	fmt.Printf("Good bye %v\n", bye)
}

func cycleNames(userNames []string, f func(string)) {
	for _, names := range userNames {
		f(names)
		fmt.Printf("Cycling through names %v\n", names)
	}
}

func cycleArea(r float64) float64 {
	return math.Pi * r * r / 10
}
