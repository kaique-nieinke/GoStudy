/* package main

import "fmt"

func main() {
	lista := []int{4, 9, 6, 7}
	fmt.Println("Lista:", lista)
	fmt.Println("Lista pos1: ", lista[0])
	fmt.Println("Lista pos2: ", lista[1])
	fmt.Println("Lista pos3: ", lista[2])
	fmt.Println("tamanho da lista: ", len(lista))

	fmt.Println("lista pos3: ", lista[5])

	lista = app
}
*/

package main

import "fmt"

func main() {
	lista := make([]int, 1)
	lista[0] = 5
	lista = append(lista, 2)
	lista = append(lista, 3)
	fmt.Printf("%T\n", lista)

	somatotal := 0
	for i := 0; i < len(lista); i++ {
		somatotal += lista[i]
	}
	fmt.Println("media:", somatotal/len(lista))

}
