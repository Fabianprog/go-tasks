package listprops

// CountValue gibt erwartet eine Liste von Zahlen und eine Zahl v.
// Liefert die Anzahl der Vorkommen von v in der Liste.
func CountValue(list []int, v int) int {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] == v {
			count++
		}

	}
	return count
}
