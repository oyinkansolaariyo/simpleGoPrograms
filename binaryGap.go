package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {

	if N < 1 || N > 2147483647 {
		return 0
	}
	n := int64(N)
	binaryRep := strconv.FormatInt(n, 2)
	var indexTracker []int
	max := 0
	for index, character := range binaryRep {
		if string(character) == "1" {
			indexTracker = append(indexTracker, index)
		}

	}

	if len(indexTracker) < 1 {
		return 0
	}

	for i := 1; i < len(indexTracker); i++ {
		diff := indexTracker[i] - indexTracker[i-1]
		if diff > max {
			max = diff -1
		}
	}
	return max

	// write your code in Go 1.4
}

func main() {
	fmt.Println(Solution(20))

}
