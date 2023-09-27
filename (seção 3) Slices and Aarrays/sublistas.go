package main

import "fmt"

func main() {
	var listatoda = []int{2, 10, 9, 4, 5, 8, 1, 3}
	var listamenorquecinco = make([]int, 0)
	for i := 0; i < len(listatoda); i++ {
		if listatoda[i] < 5 {
			listamenorquecinco = append(listamenorquecinco, listatoda[i])
		}
	}
	fmt.Println(listamenorquecinco)

}
