package main

import "fmt"

func main() {
	mensagem("mensagem x")
	mensagem("mensagem y")
}

func mensagem(mensagem string) {
	mensagem += ",bomdia!"
	fmt.Println(mensagem)
}
