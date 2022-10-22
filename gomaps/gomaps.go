package gomaps

import "fmt"

func GoMaps() {
	billionaires_networth := map[string]int{
		"Elon musk":     200,
		"Jack dorsey":   5,
		"Larry page":    145,
		"Aliko dangote": 20,
		"Mike adenuga":  10,
	}

	for networths, names := range billionaires_networth {
		fmt.Printf("Billionaire name: %v and Billionaire networth: $%v billion dollars\n", networths, names)
	}
}
