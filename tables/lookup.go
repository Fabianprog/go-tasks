package tables

// Lookup erwartet zwei Listen und einen Wert.
// Sucht den Wert in l1 und liefert den entsprechenden Wert aus l2.
// Liefert einen leeren String, falls der gesuchte Wert
// in l1 nicht vorkommt oder falls die gefundene Position nicht in l2 vorkommt.
func Lookup(l1, l2 []string, v string) string {
	// Hinweis:
	// Verwenden Sie die Funktion Find, um die Position des Werts in l1 zu ermitteln.
	// Prüfen Sie anschließend, ob die gefundene Position in l2 existiert.
	// begin:solution
	pos := Find(l1, v)
	if pos == -1 || pos >= len(l2) {
		return ""
	}
	return l2[pos]
	// end:solution
}
