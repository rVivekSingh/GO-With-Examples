package main

import "fmt"

func main() {
	numbers := []int{
		10, 2, 3, 64, 76, 44, 97, 12}

	for _, n := range numbers {
		checkNumber(n)
	}

}

func checkNumber(num int) {
	if num%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is odd")
	}
}
