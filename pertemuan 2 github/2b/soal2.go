package main

import "fmt"

func main() {
	var N int
	var bunga string
	var pita string

	fmt.Print("N: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita += bunga + " - "
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", len(string.Split(pita, " - "))-1)
}
