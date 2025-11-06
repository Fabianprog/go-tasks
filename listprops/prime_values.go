package listprops

// PrimeValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der Primzahlen in der Liste.
func PrimeValues(list []int) int {
	count := 0

	for i := 0; i < len(list); i++ {
		if Prim(list[i]) {
			count++
		}

	}
	return count

}
func Prim(p int) bool {
	div := 0
	for i := 1; i <= p; i++ {
		if p%i == 0 {
			div++
		}

	}
	return div == 2

}
