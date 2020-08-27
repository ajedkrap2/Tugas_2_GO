package main

import (
    "fmt"
)

func main() {
    var angka = 255
    switch angka%2 {
    case 0: fmt.Printf("%d adalah bilangan genap", angka)
    default: fmt.Printf("%d adalah bilangan ganjil", angka)
    }
    fmt.Println()
}
