package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("masukan input string:")
	fmt.Scanln(&satu)
	fmt.Print("masukan input string:")
	fmt.Scanln(&dua)
	fmt.Print("masukan input string:")
	fmt.Scanln(&tiga)
	fmt.Println("string satu: " + satu + "" + dua + "" + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("output akhir: " + satu + "" + dua + "" + tiga)

}
