package main

import "fmt"

func Fibonacci(number int) int {
	a := 0
	b := 1
	c := 0

	if number < 2 {
		return number
	}

	for i := 1; i < number; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func main() {
	fmt.Println(Fibonacci(0))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(9))
	fmt.Println(Fibonacci(10))
	fmt.Println(Fibonacci(12))
}
