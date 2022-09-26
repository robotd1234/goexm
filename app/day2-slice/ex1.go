package main

import "fmt"

func main() {

	slice1 := []int{0, 1, 2}
	slice2 := make([]int, 3, 6)
	slice3 := make([]int, 2, 3)
	slice3 = []int{0, 1, 2, 3}
	slice4 := make([]int, 2, 4)
	slice4 = slice3[1:3:4]

	fmt.Println(len(slice1))
	fmt.Println(len(slice2))
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(slice4)
	fmt.Println(len(slice4))

	// str1 := "hello world"
	// str2 := "the go web"
	// str3 := str1 + " " + str2
	// fmt.Printf("the string is %s\n", str3)

	// for key, value := range str3 {
	// 	fmt.Printf("%d %U %c \n", key, value, value)
	// }

	// var array []int
	// array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	// fmt.Println(array)

}
