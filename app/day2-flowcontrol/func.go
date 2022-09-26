package main

import "fmt"

func main() {

	// arr := []int{10, 2, 3, 4, 3, 6, 7, 8, 9}
	// var m int
	// m = min(arr)
	// println(m)

	num1 := 6
	num2 := 10
	fmt.Printf("num1 before change:%d\n", num1)
	fmt.Printf("num2 before change:%d\n", num2)
	num1, num2 = exchange(num1, num2)
	fmt.Printf("num1 after change:%d\n", num1)
	fmt.Printf("num2 after change:%d\n", num2)
}

func min(arr []int) (m int) {
	m = arr[0]
	for _, v := range arr {
		if v < m {
			m = v
		}
	}
	return
}

func exchange(x, y int) (int, int) {

	var temp int
	temp = x
	x = y
	y = temp
	return x, y
}
