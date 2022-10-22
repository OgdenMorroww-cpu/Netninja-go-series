package loops

import "fmt"

func Loops() {
	// index := 0

	// for index < 1000_000_000_000_000_000 {
	// 	fmt.Printf("Index: %v\n", index)
	// 	index++
	// }

	// for i := 0; i <= 120; i++ {
	// 	fmt.Printf("Numbers %v\n", i)
	// }

	black_adam_cast := []string{"The rock", "Pierce brosnan", "Aldis Hoge", "Henry cavill"}

	// for i := 0; i < len(black_adam_cast); i++ {
	// 	fmt.Printf("Cast Of Black Adam: %v\n", black_adam_cast[i])
	// }

	for index, cast := range black_adam_cast {
		fmt.Printf("The position at index %v is %v\n", index, cast)
	}
}
