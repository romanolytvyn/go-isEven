package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range nums {
		even := isEven(n)
		if even {
			fmt.Println("Number", n, "is even")
		} else {
			fmt.Println("Number", n, "is odd")
		}
	}
}

func isEven(num int) bool {
	return (num % 2) == 0
}
