package main

import "fmt"

func main() {
	var r int
	fmt.Print("Enter number of rows")
	fmt.Scan(&r)

	for i := 0; i < r; i++ {
		for s := 0; s < (r - i); s++ {
			fmt.Print("  ")
		}

		for a := 0; a < (2*i - 1); a++ {
			fmt.Print("* ")
		}

		fmt.Println("")

	}
}
