/*package main

import "fmt"

func main() {
	salario := 5500.00

	if salario < 1000.00 {
		salario -= (salario * 0.08)

	} else if salario <= 1200.00 {
		salario -= (salario * 0.10)

	} else if salario > 5000.00 {
		salario -= (salario * 0.08)

	} else {
		salario -= (salario * 0.15)
	}

	fmt.Println("salario final:", salario)
}
*/

package main

import "fmt"

func main() {
	salario := 900.00
	tipo := "PJ"

	if salario < 1000.00 && tipo == "CLT" {
		salario -= (salario * 0.08)

	} else if salario <= 1000.00 && tipo == "PJ" {
		salario -= (salario * 0.05)

	} else {
		salario -= (salario * 0.15)
	}

	fmt.Println("salario final:", salario)
}
