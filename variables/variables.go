package variables

import "fmt"

var areaOfFocus = "Physics"

func Variables() {
	var name string = "mario"
	var secondName string = "luigi"
	var thirdName string
	fmt.Println(name)
	fmt.Println(secondName)
	fmt.Println(thirdName)
	name = "Nathan Rosen"
	secondName = "Albert Einstein"
	thirdName = "Werner Heisenberg"
	fmt.Println(name, secondName, thirdName)
	recentPhysics := "Richard Feynman"
	fmt.Println(recentPhysics)
	fmt.Println(areaOfFocus)
	areaOfFocus = "Biological Computation"
	fmt.Println(areaOfFocus)
	integers()
	myFloat()
}

func integers() {
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)
	var num1 int8 = 45
	num1 = -128
	var num2 int16 = 4560
	var num3 uint8 = 120
	fmt.Println(num1, num2, num3)
}

func myFloat() {
	var scorePoints float32 = 458.5
	var scorePointsTwo float64 = 8888888456783
	scorePointThree := 1.5
	fmt.Println(scorePoints, scorePointsTwo, scorePointThree)
}
