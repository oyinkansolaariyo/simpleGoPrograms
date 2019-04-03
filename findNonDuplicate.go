package main

import "fmt"

func Solution(A []int) int {
	// write your code in Go 1.4
	myMap := make(map[int]int)
	for i := 0; i < len(A); i++ {
		if myMap[A[i]] == 0 {
			myMap[A[i]] = 1
		} else {
			myMap[A[i]] += 1
		}
	}
	unpaired := 0
	for k, v := range myMap {
		if v == 1 {
			unpaired = k
			break

		}
	}
	return unpaired
}

func main() {
	fmt.Println(Solution([]int{9, 3, 9, 3, 9, 7, 9}))
}
