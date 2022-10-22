package arrays

import "fmt"

func Slices() {
	scores := []int{10, 45, 50, 120, 100}
	fmt.Printf("Scores: %v\n", scores)
	scores = append(scores, 120, 45, 678, 1890, 45)
	fmt.Printf("Updated scores: %v\n", scores)

	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	for _, planet := range planets {
		// planets = append(planets, "Proxima B", "Kepler B-124")
		fmt.Printf("Planets: %v\n", planet)
	}
}
