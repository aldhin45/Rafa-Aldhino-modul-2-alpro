package main

import "fmt"

func main() {
	var nam float64
	var nilaiAkhir string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)
	if nam > 80 {
		nilaiAkhir = "A"
	} else if nam > 72.5 {
		nilaiAkhir = "AB"
	} else if nam > 65 {
		nilaiAkhir = "B"
	} else if nam > 57.5 {
		nilaiAkhir = "BC"
	} else if nam > 50 {
		nilaiAkhir = "C"
	} else if nam > 40 {
		nilaiAkhir = "D"
	} else {
		nilaiAkhir = "E"
	}
	fmt.Println("Nilai mata kuliah: ", nilaiAkhir)
}
