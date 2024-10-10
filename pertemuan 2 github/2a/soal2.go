package main

import "fmt"

func isLeapYear(year int) bool {
    // A leap year is divisible by 400 or divisible by 4 but not divisible by 100
    return (year % 400 == 0) || (year % 4 == 0 && year % 100 != 0)
}

func main() {
    var year int

    fmt.Print("Tahun: ")
    fmt.Scan(&year)

    if isLeapYear(year) {
        fmt.Println("Kabisat: true")
    } else {
        fmt.Println("Kabisat: false")
    }
}