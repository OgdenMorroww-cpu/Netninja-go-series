package standard

import (
	"fmt"
	"sort"
	"strings"
)

func GoPackage() {
	greetings := "hello there guys"
	fmt.Println(strings.Contains(greetings, "there"))
	fmt.Println(strings.ReplaceAll(greetings, "there", "my"))

	userName := "Shedrack"
	fmt.Println(strings.Index(userName, "Name"))

	fmt.Println(strings.Split(userName, " "))

	prices := []int{120, 45, 10, 23, 890, 1000, 1200}

	sort.Ints(prices)
	fmt.Println(prices)

	index := sort.SearchInts(prices, 1200)
	fmt.Println(index)

	cars := []string{"Bugatti", "BMW", "Audi", "Alfa Romeo", "Ferrari"}
	sort.Strings(cars)
	fmt.Println(cars)

	getCars := sort.SearchStrings(cars, "BMW")
	fmt.Printf("Car Position Index: %v\n", getCars)
}
