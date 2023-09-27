/*package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"sp": 1000000, "cg": 90000}
	fmt.Println(m)
}
*/

package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["sp"] = 100000000
	m["cg"] = 900000

	fmt.Println(m)
}
