package reservations

// ContainsOverlap erwartet eine Liste von Strings und zwei Paare von jeweils zwei Strings.
// Liefert true, falls die String-Paare sich in der Liste überlappen,
// also falls z.B. t1 zwischen s1 und s2 liegt.
// Anmerkung: Die Reihenfolge der Strings im Paar ist hier nicht relevant.
// Anmerkung: An den Grenzen (also s1 == t1 oder s2 == t2) liegt kein Überlappen vor.
func ContainsOverlap(list []string, s1, s2, t1, t2 string) bool {
	eins := false
	zwei := false
	einseins := false
	zweizwei := false
	var higher1 int
	var higher2 int
	var lower1 int
	var lower2 int
	var ss1 int
	var ss2 int
	var tt1 int
	var tt2 int

	for i := 0; i < len(list); i++ {
		if list[i] == s1 {
			eins = true
			ss1 = i
		}
		if list[i] == s2 {
			zwei = true
			ss2 = i
		}
		if list[i] == t1 {
			einseins = true
			tt1 = i
		}
		if list[i] == t2 {
			zweizwei = true
			tt2 = i
		}

	}
	if eins == false ||
		zwei == false ||
		einseins == false ||
		zweizwei == false {
		return false
	}
	if ss1 > ss2 {
		higher1 = ss1
	} else {
		higher1 = ss2
	}
	if tt1 > tt2 {
		higher2 = tt1
	} else {
		higher2 = tt2
	}
	if ss1 <= ss2 {
		lower1 = ss1
	} else {
		lower1 = ss2
	}
	if tt1 <= tt2 {
		lower2 = tt1
	} else {
		lower2 = tt2
	}
	if lower1 && higher1 

}
