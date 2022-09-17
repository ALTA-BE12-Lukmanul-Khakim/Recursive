package main

import "fmt"

func MaxSequence(arr []int) int {
	var res int
	//min := 0
	sSum := 0

	for i := 0; i < len(arr); i++ {
		sSum = sSum + arr[i]
		if res < sSum {
			res = sSum
		}
		if sSum < 0 {
			sSum = 0
		}
	}
	return res
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
