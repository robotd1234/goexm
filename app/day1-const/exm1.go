// Training excercises for golang. For learning uses.
// by bobo

package main

import (
	"fmt"
)

func main() {

	// decleration block

	type with int
	const (
		a with = iota
		b
		c
		d
		e
		f
		g
	)

	// Logic block

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
