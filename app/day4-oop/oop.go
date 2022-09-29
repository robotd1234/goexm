package main

import "fmt"

type Voice interface {
	cry() string
	laught() string
	interested() string
}

type Profile interface {
	height() float32
	weight() float32
	gender() string
}

type People struct {
	name   string
	income []int
	id     map[string]string
}

type Animal struct {
	name string
	id   map[string]string
}

func (aoao *People) cry() string {
	fmt.Printf("this is people cry: %v\n", aoao.name)
	s := ""
	return s
<<<<<<< HEAD
=======
>>>>>>> origin/branch1
}

func main() {

}
