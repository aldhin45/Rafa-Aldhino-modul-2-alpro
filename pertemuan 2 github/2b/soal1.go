package main

import "fmt"

func main() {
    var warna1, warna2, warna3, warna4 string
    var berhasil bool = true

    for i := 1; i <= 5; i++ {
        fmt.Printf("Percobaan %d: ", i)
        fmt.Scan(&warna1, &warna2, &warna3, &warna4)

        if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
            berhasil = false
            break
        }
    }

    if berhasil {
        fmt.Println("BERHASIL: true")
    } else {
        fmt.Println("BERHASIL: false")
    }
}