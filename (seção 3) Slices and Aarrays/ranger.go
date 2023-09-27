package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["sp"] = 100000000
	m["cg"] = 900000
	/*delete(m,"XP") se usa quando é preciso deletar essa determinada chave*/

	for chave, valor := range m {
		fmt.Println("cidade", chave, "H", valor)
	}
}
