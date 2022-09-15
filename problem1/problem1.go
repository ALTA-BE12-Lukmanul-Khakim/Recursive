package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeX(number int) int {
	//	var hasil int
	k := 0
	//	var list [number] int
	if number < 3 {
		return number + 1
	}

	i := 2
	for ; ; i++ {
		if isPrime(i) == true {
			k++
		}
		if k == number {
			break
		}
	}
	return i
}

func main() {
	fmt.Println(PrimeX(1))
	fmt.Println(PrimeX(5))
	fmt.Println(PrimeX(8))
	fmt.Println(PrimeX(9))
	fmt.Println(PrimeX(10))
}
