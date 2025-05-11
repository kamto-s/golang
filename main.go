package main

import "fmt"

func main() {
	// operator logika
	/*
		- AND (&&)
		- OR (||)
		- NOT (!)
	*/

	fmt.Printf("true && false: %t \n", 1 == 1 && 3 == 2)
	fmt.Printf("true || false: %t \n", 1 == 1 || 3 == 2)
	fmt.Printf("!true: %t \n", !(1 == 3))
}
