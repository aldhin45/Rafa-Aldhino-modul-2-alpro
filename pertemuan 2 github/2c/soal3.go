package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)
	fmt.Print("Faktor : ")
	for i := 1; i < bilangan+1; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
