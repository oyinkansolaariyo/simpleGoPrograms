package main

import "fmt"

func main(){
   fmt.Println(Solution([]int{4,1,3,2}))
}

func Solution(A []int) int {
    // write your code in Go 1.4
    permutation := 1
    if len(A) < 1 || len(A) > 100000 {
        return 0
    }
    myMap := make(map[int]int)
    for i := 0; i < len(A); i++ {
		if myMap[A[i]] == 0 {
			myMap[A[i]] = 1
		} else {
			myMap[A[i]] += 1
		}
	}
	
	
	for j:= 1; j <= len(A); j++ {
	    if myMap[j] == 0 {
	        permutation = 0
	        break
	    }
	    
	}
    return permutation
    
}
