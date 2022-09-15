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

func prismaSegiEmpat(wide, high, start int) string {
	//var k int
	var res string
	for i := 0; i <= high; i++ {
		for j := 0; j <= wide; j++ {
			start++
			if isPrime(start) == true {
				res = res + fmt.Sprintf("%d", start)
			}

		}
		res = res + "\n"
	}
	return res
}

func main() {
	fmt.Println(prismaSegiEmpat(2, 3, 13))
	//fmt.Println(prismaSegiEmpat(2, 3, 13))
}
