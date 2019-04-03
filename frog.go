package main

import "fmt"
//import "math"

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	if X < 1 && X > 1000000000 {
		return -1
	}
	if Y < 1 && Y > 1000000000 {
		return -1
	}

	if Y < X {
		return -1
	}

	if D < 1 && D > 1000000000 {
		return -1
	}

	diff := Y - X
	fmt.Println(diff % D)
	if diff % D == 0 {
	  return diff/D	
	} else {
	  return (diff/D) + 1
	}

}

func main() {
	fmt.Println(Solution(10, 85, 30))

}
