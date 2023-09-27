package main

import "fmt"

func exibirCombinacoes() {
    for i := 0; i <= 99; i++ {
        for j := i + 1; j <= 99; j++ {
            fmt.Printf("%02d %02d, ", i, j)
        }
    }
    fmt.Println()
}

func main() {
    fmt.Println("Combinações de dois números de dois dígitos em ordem ascendente:")
    exibirCombinacoes()
}
