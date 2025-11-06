package listprops

// ContainsValue gibt erwartet eine Liste von Zahlen und eine Zahl v.
// Liefert true, falls v in der Liste enthalten ist, sonst false.
func ContainsValue(list []int, v int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == v {
			return true
		}

	}
	return false
}
