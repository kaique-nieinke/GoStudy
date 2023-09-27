package main

import "fmt"

func main() {
	for numbase := 1; numbase <= 10; numbase++ {

		for multiplicado := 1; multiplicado <= 10; multiplicado++ {

			fmt.Println(numbase, " x ", multiplicado, " = ", numbase*multiplicado)

		}
	}
}
