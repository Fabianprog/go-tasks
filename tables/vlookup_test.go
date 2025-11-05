package tables

import "fmt"

func ExampleVLookup() {
	table := [][]string{
		{"Haus", "house"},
		{"Apfel", "apple"},
		{"Holz", "wood"},
		{"Wald", "forest"},
		{"Auto", "car"},
		{"Vorlesung", "lecture"},
		{"Fahrrad", "bicycle"},
	}

	fmt.Println(VLookup(table, "Haus", 1))
	fmt.Println(VLookup(table, "Apfel", 1))
	fmt.Println(VLookup(table, "Vorlesung", 1))
	fmt.Println(VLookup(table, "Birne", 1))

	// Output:
	// house
	// apple
	// lecture
	//
}
