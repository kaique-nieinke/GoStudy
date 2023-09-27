package main

import "fmt"

/*func main()  {
	x := 5
	y := &x
	*y = 10
	fmt.Println("main-----------------------")
	fmt.Println(x, *y)
	fmt.Println(&x, y)
	imprimirValores(&x, y)
}

func imprimirValores (x *int, y *int) {
	fmt.Println("imprimirValores---------------")
	fmt.Println(x,y)
}
*/
func main()  {
	x := 5
	y := &x
	*y = 10
	imprimirValores(&x,y)
	fmt.Println(x, *y)
	fmt.Println(&x, y)
}

func imprimirValores (x *int, y *int) {
	*x = 20
}