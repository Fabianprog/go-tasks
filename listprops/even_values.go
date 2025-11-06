package listprops

// EvenValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der geraden Zahlen in der Liste.
func EvenValues(list []int) int {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i]%2 == 0 {
			count++
		}

	}
	return count

}
