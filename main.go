package main

import (
	_ "fmt"
	_ "strconv"
)

func main() {

	var stücke = [10]string{"Hose", "Pullover", "T-Shirt", "Kleid", "Jacke", "Sakko", "Schuhe", "Krawatte", "Hut", "Schal"}

	println("Kleidungsstücke:")

	for i := 0; i < 10; i++ {
		println(stücke[i])
	}

	println("_________________________________________")
	println("Kleidungsstücke im Schrank:")

	var schrank = [10]string{"", "", "", "", "", "", "", "", "", ""}

	for i := 0; i < 10; i++ {
		schrank[i] = stücke[i]

	}
	for i := 0; i < 10; i++ {

		println(i, schrank[i])

	}
	println("_________________________________________")

	var outfit = [5]string{"", "", "", "", ""}

	outfit[0] = schrank[0]
	outfit[1] = schrank[1]
	outfit[2] = schrank[2]
	outfit[3] = schrank[6]
	outfit[4] = schrank[8]

	println("Outfit:")
	for i := 0; i < 5; i++ {
		println(outfit[i])

	}

}
