package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println("result is ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("result is ", res)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}