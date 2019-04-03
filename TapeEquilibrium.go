package main

// you can also use imports, for example:
// import "fmt"
import "math"
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4

	rightSum := 0
	leftSum := 0
	if len(A) < 2 || len(A) > 100000 {
		return 0
	}
	tape := len(A) - 1
	var minDifference float64 = 0
	for i := 0; i < len(A); i++ {
		rightSum += A[i]
	}
	//this works for both large and small inputs
	for i := 0; i < tape; i++ {
		leftSum += A[i]
		rightSum -= A[i]
		tempDiff := math.Abs(float64(rightSum - leftSum))

		if i == 0 {
			minDifference = tempDiff
		} else {
			minDifference = math.Min(tempDiff, minDifference)
		}
		fmt.Println(minDifference)

	}
	//this works for only small and medium inputs
	/**for i := 0; i < tape; i++ {
		leftSum := 0
		rightSum := 0
		if i == 0 {
			leftSum = A[0]

		} else {
			for l := 0; l <= i; l++ {
				leftSum += A[l]
			}
		}
		for r := (i + 1); r < len(A); r++ {
			rightSum += A[r]
		}
		fmt.Println(rightSum, leftSum)
		tempDiff := math.Abs(float64(rightSum - leftSum))

		if i == 0 {
			minDifference = tempDiff
		} else {
			minDifference = math.Min(tempDiff, minDifference)
		}
		fmt.Println(minDifference)

	}*/
	return int(minDifference)

}

func main() {
	fmt.Println(Solution([]int{1, 2, 3, 4, 2}))
}
