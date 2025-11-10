package reservations

// ReservationsOk erwartet eine Route (Liste von Stationen (Strings))
// eine Liste von Reservierungen (jeweils ein Paar von Strings).
// Liefert true, falls alle Reservierungen in der Route erfüllt werden können,
// also falls für jede Reservierung beide Stationen in der richtigen Reihenfolge
// in der Route vorkommen und falls es keine sich überlappenden Reservierungen gibt.

func ReservationsOk(route []string, reservations [][2]string) bool {
	for i := 0; i < len(reservations); i++ {
		start := reservations[i][0]
		end := reservations[i][1]

		if !ContainsBoth(route, start, end) {
			return false
		}
		if !ContainsInOrder(route, start, end) {
			return false
		}
		for s := 0; s < len(reservations); s++ {
			startcomp := reservations[s][0]
			endcomp := reservations[s][1]
			if i == s {
				continue
			}
			if ContainsOverlap(route, start, end, startcomp, endcomp) {
				return false
			}

		}

	}
	return true
}
