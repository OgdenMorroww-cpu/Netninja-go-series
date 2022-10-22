package standard

import (
	"fmt"
	"sort"
)

func SortCars() {
	usedCars := []string{"Toyota", "Mercedes", "Bentley", "BMW", "Honda", "Tesla"}
	sortCars := sort.SearchStrings(usedCars, "Bentley")

	fmt.Printf("Index: %v\n", sortCars)
}
