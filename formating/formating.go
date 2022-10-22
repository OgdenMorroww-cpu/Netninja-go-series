package formating

import "fmt"

var (
	age      = 45
	userName = "Shedrack"
	country  = "Cambodia"
	job      = "Computer programmer"
	points   = 45.8
)

func Formating() {
	fmt.Print("Hello, ")
	fmt.Print("World \n")

	//Println
	fmt.Println("Hello kryptonians")

	fmt.Println("My UserName is", userName)
	fmt.Println("My Age is", age)

	//Printf
	fmt.Printf("My age is %v and my name is %v\n", age, userName)
	// fmt.Printf("Hey you are %q and your name is %q\n", age, userName)
	fmt.Printf("You are from: %q and you are a: %q\n", country, job)
	fmt.Printf("Age is of type %T and Name is of type %T\n", age, userName)
	fmt.Printf("you scored %0.1f points so far\n", points)
	var str = fmt.Sprintf("my age is %v and my userName is %v\n", age, userName)
	fmt.Println("The saved String is:", str)
}
