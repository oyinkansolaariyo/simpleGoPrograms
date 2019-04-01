package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Enter the number")
	var number int
	fmt.Scan(&number)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("A panic has occured")
		}
	}()

	fmt.Println("The factorial is", factorial(number))

}
