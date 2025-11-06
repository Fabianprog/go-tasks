package listprops

// OddValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen in der Liste.
func OddValues(list []int) int {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i]%2 != 0 {
			count++
		}

	}
	return count

}
