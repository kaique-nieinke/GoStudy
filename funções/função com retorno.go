package main

import "fmt"

func main() {
	resultado := soma(1, 2)
	fmt.Println(resultado)
}

func soma(numero1 int, numero2 int) int {
	resultado := numero1 + numero2
	return resultado
}
