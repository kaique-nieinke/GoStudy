/*package main

import "fmt"

func main() {
	var ehcarro bool = false
	var valorDoAutomovel = 1000.00

	if ehcarro {
		valorDoAutomovel = valorDoAutomovel + 55.50
	} else {
		valorDoAutomovel += 20.50
	}

	fmt.Println("valor final:", valorDoAutomovel)
}
*/

package main

import "fmt"

func main() {
	salario := 1202.00

	if salario < 1000.00 {
		salario -= (salario * 0.08)
	} else if salario <= 1200.00 {
		salario -= (salario * 0.10)
	} else {
		salario -= (salario * 0.10)
	}

	fmt.Println("salario final:", salario)
}
