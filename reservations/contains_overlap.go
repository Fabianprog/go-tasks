package reservations

// ContainsOverlap erwartet eine Liste von Strings und zwei Paare von jeweils zwei Strings.
// Liefert true, falls die String-Paare sich in der Liste überlappen,
// also falls z.B. t1 zwischen s1 und s2 liegt.
// Anmerkung: Die Reihenfolge der Strings im Paar ist hier nicht relevant.
// Anmerkung: An den Grenzen (also s1 == t1 oder s2 == t2) liegt kein Überlappen vor.

func ContainsOverlap(list []string, s1, s2, t1, t2 string) bool {
	var ss1, ss2, tt1, tt2 int
	found := 0

	for i := 0; i < len(list); i++ {
		if list[i] == s1 {
			ss1 = i
			found++
		}
		if list[i] == s2 {
			ss2 = i
			found++
		}
		if list[i] == t1 {
			tt1 = i
			found++
		}
		if list[i] == t2 {
			tt2 = i
			found++
		}
	}
	if found < 4 {
		return false
	}

	// Intervallgrenzen berechnen
	lower1 := min(ss1, ss2)
	higher1 := max(ss1, ss2)
	lower2 := min(tt1, tt2)
	higher2 := max(tt1, tt2)

	// Kein Überlappen, wenn eines Intervall komplett vor dem anderen liegt
	if higher1 <= lower2 || higher2 <= lower1 {
		return false
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
