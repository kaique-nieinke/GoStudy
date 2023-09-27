package main

import "fmt"

func main() {
	salario := 850.00
	var salariMaisBonus float64

	salariMaisBonus = salario

	if salario < 1000 {
		salariMaisBonus = (salariMaisBonus + 100)
	}

	fmt.Println("Salario: ", salariMaisBonus)
}
