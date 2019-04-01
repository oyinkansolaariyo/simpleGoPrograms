package main

import "fmt"

//0,1,1,2,3,5,8,13,21
func fibonacci(number int) int {
	if number < 2 {
		return number
	}

	return fibonacci(number-1) + fibonacci(number-2)

}

func main() {
	var n int
	fmt.Print("Enter number of terms")
	fmt.Scan(&n)
	fmt.Print("Fibonacci series")

	for i := 1; i <= n; i++ {
		fmt.Println(fibonacci(i))

	}

}
