package reservations

// ContainsInOrder erwartet eine Liste von Strings und zwei Strings s1 und s2.
// Liefert true, falls sowohl s1 als auch s2 in der Liste enthalten sind, und s1 vor s2 kommt.
func ContainsInOrder(list []string, s1 string, s2 string) bool {
	eins := false
	zwei := false
	var ort1 int
	var ort2 int
	for i := 0; i < len(list); i++ {
		if list[i] == s1 {
			eins = true
			ort1 = i
		}
		if list[i] == s2 {
			zwei = true
			ort2 = i
		}
	}
	if eins == false || zwei == false {
		return false
	}
	if ort1 < ort2 {
		return true
	}
	return false

}
