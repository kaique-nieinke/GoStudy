package main

import "fmt"

func init() {
	fmt.Println("função init2")
}

func main() {
	soma, subtracao, divisao, multiplicacao := operação(3, 5)
	fmt.Println(soma, subtracao, divisao, multiplicacao)
}

func operação(numero1 int, numero2 int) (soma int, subtracao int, divisao int, multiplicacao int) {
	soma = numero1 + numero2
	subtracao = numero1 - numero2
	divisao = numero1 / numero2
	multiplicacao = numero1 * numero2
	return
}

func init() {
	fmt.Println("função init1") // (func init) serve para inicializar a programa, independente de onde esteja posicionada, ela sempre vai ser executada primeiro//
}
