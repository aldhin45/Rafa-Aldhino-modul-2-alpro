package main

import "fmt"

func main() {
	var gram, beratkg, beratGram, biayakg, biayagram int
	fmt.Print("Berat parsel (gram) : ")
	fmt.Scan(&gram)
	beratkg = gram / 1000
	beratGram = gram % 1000
	biayakg = beratkg * 10000
	if beratGram > 500 {
		biayagram = beratGram * 5
	} else {
		biayagram = beratGram * 15
	}
	fmt.Printf("Detail berat: %v kg + %v gr\n", beratkg, beratGram)
	fmt.Printf("Detail biaya: Rp. %v + Rp. %v \n", biayakg, biayagram)
	fmt.Println("Total biaya: Rp. ", biayagram+biayakg)
}
