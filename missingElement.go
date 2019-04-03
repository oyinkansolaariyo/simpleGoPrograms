package main

import "fmt"
import "sort"

func Solution(A []int) int {
	// write your code in Go 1.4
	//it fails for [1,2]
	if len(A) > 1000000 {
        return 1
    }
    
    if len(A) == 0 {
        return 1
    }
    
    if len(A)  == 1 {
        if  A[0] > 1 {
            return A[0]-1
        } else {
            return A[0] + 1
        }
    }
    sort.Ints(A)
	var missing int = 0
	for i := 0; i < len(A); i++ {
		if (A[i+1] - A[i]) > 1 {
			missing = A[i+1] - 1
			break
		}
	}
	return missing
}

func main() {
	fmt.Println(Solution([]int{1,2}))
}
