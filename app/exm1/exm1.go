// Training excercises for golang. For learning uses.
// by bobo

package main

import (
	"fmt"
)

func main() {

	// decleration block

	const name string = "Tom"
	const name_1, name_2 string = "Tom", "Jay"
	const name_3, age_1 = "Tom", 30
	const age = 30

	// Logic block

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(name_1, name_2)
	fmt.Println(name_3, age_1)
}
