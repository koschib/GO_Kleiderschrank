package main

import (
	_ "fmt"
	_ "strconv"
)

func main() {

	// Array erstellen mit 10 Kleidungsstücken
	var stücke = [10]string{"Hose", "Pullover", "T-Shirt", "Kleid", "Jacke", "Sakko", "Schuhe", "Krawatte", "Hut", "Schal"}

	// Array durchlaufen
	println("Kleidungsstücke:")
	// Array ausgeben
	for i := 0; i < 10; i++ {
		println(stücke[i])
	}

	println("_________________________________________")
	println("Kleidungsstücke im Schrank:")
	// Array erstllen Schrank mit 10 Plätzen
	var schrank = [10]string{"", "", "", "", "", "", "", "", "", ""}
	// elemente von clothes in schrank kopieren
	for i := 0; i < 10; i++ {
		schrank[i] = stücke[i]

	} // Array ausgeben
	for i := 0; i < 10; i++ {
		// Ausgabe Index und Inhalt
		println(i, schrank[i])

	}
	println("_________________________________________")

	// Outfit erstellen
	var outfit = [5]string{"", "", "", "", ""}
	// Outfit zusammenstellen
	outfit[0] = schrank[0]
	outfit[1] = schrank[1]
	outfit[2] = schrank[2]
	outfit[3] = schrank[6]
	outfit[4] = schrank[8]

	// Ausgabe Outfit
	println("Outfit:")
	for i := 0; i < 5; i++ {
		println(outfit[i])

	}

}
