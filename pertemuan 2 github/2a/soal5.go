package main

import "fmt"

func main() {
    var integers [5]int
    var characters [3]byte

    // Read 5 integers
    fmt.Println("Masukkan 5 buah data integer (32-127):")
    for i := 0; i < 5; i++ {
        fmt.Scan(&integers[i])
    }

    // Read 3 characters
    fmt.Println("Masukkan 3 buah karakter:")
    for i := 0; i < 3; i++ {
        fmt.Scanf("%c", &characters[i])
    }

    // Print characters from integers
    fmt.Println("Karakter dari data integer:")
    for _, integer := range integers {
        fmt.Printf("%c", integer)
    }
    fmt.Println()

    // Print characters after incrementing ASCII values
    fmt.Println("Karakter setelah diinkremen:")
    for _, character := range characters {
        fmt.Printf("%c", character+1)
    }
    fmt.Println()
}