/*package main

import "fmt"

func main() {
	salario := 1000.00
	var salariMaisBonus float64

	salariMaisBonus = salario

	if salario <= 1000 {
		salariMaisBonus = (salariMaisBonus + 100)
	}

	fmt.Println("Salario: ", salariMaisBonus)
}*/

package main

import "fmt"

func main() {
	var ehcarro bool = true
	var valorDoAutomovel = 1000.00

	if ehcarro {
		valorDoAutomovel = valorDoAutomovel + 55.50
	}

	fmt.Println("valor final:", valorDoAutomovel)
}

/*
	<   menor
	>	maior
	>=  maior igual
	<=  menor igual
*/
