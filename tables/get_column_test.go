package tables

import "fmt"

func ExampleGetColumn() {
	table := [][]string{
		{"A1", "B1", "C1", "D1"},
		{"A2", "B2", "C2"},
		{"A3", "B3", "C3", "D3"},
	}

	fmt.Println(GetColumn(table, 0))
	fmt.Println(GetColumn(table, 1))
	fmt.Println(GetColumn(table, 2))
	fmt.Println(GetColumn(table, 3))
	fmt.Println(GetColumn(table, 4))

	// Output:
	// [A1 A2 A3]
	// [B1 B2 B3]
	// [C1 C2 C3]
	// [D1  D3]
	// [  ]
}
