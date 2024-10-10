package main

import (
        "fmt"
        "math"
)

func main() {
        var beratKiri, beratKanan float64

        for {
                fmt.Print("Masukkan berat belanjaan di kedua kantong (kg): ")
                fmt.Scan(&beratKiri, &beratKanan)

                selisih := math.Abs(beratKiri - beratKanan)

                if selisih >= 9 || beratKiri >= 9 || beratKanan >= 9 {
                        fmt.Println("Sepeda motor Pak Andi akan oleng: true")
                        break
                } else {
                        fmt.Println("Sepeda motor Pak Andi akan oleng: false")
                }
        }

        fmt.Println("Proses selesai.")
}