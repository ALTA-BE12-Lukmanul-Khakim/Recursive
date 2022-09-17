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

// func listPrime (n int) [] int{

// }

func PrismaSegiEmpat(wide, high, start int) string {
	var res string
	var k, sum int
	primeArr := []int{}
	var n = wide * high
	for n > 0 {
		start++
		if isPrime(start) == true {
			primeArr = append(primeArr, start)
			n--
		}
	}
	for i := 0; i < high; i++ {
		for j := 0; j < wide; j++ {
			k++
			res += fmt.Sprintf("%d ", primeArr[k-1])
			sum += primeArr[k-1]
		}
		res = res + "\n"
	}
	return res + fmt.Sprintf("%d", sum)
}

func main() {
	fmt.Println(PrismaSegiEmpat(2, 3, 13))
	fmt.Println(PrismaSegiEmpat(5, 2, 1))
}
