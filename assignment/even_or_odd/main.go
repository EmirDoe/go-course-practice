package main

import "fmt"

var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {

	for i := range numbers {
		if i%2 != 0 {
			fmt.Println("This number: ", i, " is odd")
		}
		if i%2 == 0 {
			fmt.Println("This number ", i, " is even")
		}
	}
}
