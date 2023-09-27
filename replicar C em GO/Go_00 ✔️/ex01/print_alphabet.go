package main

import "fmt"

func mostrarAlfabetoMinusculo() {
    alfabeto := ""
    for i := 'a'; i <= 'z'; i++ {
        alfabeto += string(i)
    }
    fmt.Println(alfabeto)
}

func main() {
    // Chame a função para exibir o alfabeto
    mostrarAlfabetoMinusculo()
}
 