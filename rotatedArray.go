package main

import "fmt"

func Solution(A []int, K int) []int {
	expectedLength := len(A)
	if expectedLength < 1 {
	  return A
	}
	rotatedArray := make([]int, expectedLength)
	for i := 0; i < expectedLength; i++ {
	if (i+ K) >= expectedLength {
	   rotatedArray[i+K-expectedLength ] = A[i]
	}  else {
		rotatedArray[i+K] = A[i]
	}
		
	}
	// write your code in Go 1.4
	return rotatedArray
}

func main() {
	fmt.Println(Solution([]int{3, 8, 9, 7, 6}, 3))
}
