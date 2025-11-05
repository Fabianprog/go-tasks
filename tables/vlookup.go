package tables

// VLookup erwartet ein zweidimensionales Array (Tabelle), einen Suchwert und eine Spaltennummer.
// Sucht den Wert in der ersten Spalte der Tabelle und liefert den entsprechenden Wert aus der angegebenen Spalte.
// Liefert einen leeren String, falls der gesuchte Wert nicht in der ersten Spalte vorkommt
// oder falls die gefundene Position nicht in der angegebenen Spalte vorkommt.
func VLookup(table [][]string, v string, col int) string {
	// Hinweis:
	// Verwenden Sie die Funktion GetColumn, um die erste Spalte und die angegebene Spalte zu extrahieren.
	// Verwenden Sie anschließend die Funktion Lookup, um den gesuchten Wert zu finden.
	// begin:solution
	firstCol := GetColumn(table, 0)
	return Lookup(firstCol, GetColumn(table, col), v)
	// end:solution
}
