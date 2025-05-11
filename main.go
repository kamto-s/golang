package main

import "fmt"

func main() {
	fmt.Println("==========================")
	i := 8
	for i <= 10 {
		fmt.Printf("Nilai ke-%d \n", i)
		i++
	}

	fmt.Println("==========================")

	for i := 1; i <= 10; i++ {
		fmt.Printf("Nilai ke-%v \n", i)
	}

	fmt.Println("==========================")

	for i := 1; i <= 10; i += 2 {
		fmt.Printf("Nilai ke-%v \n", i)
	}

	fmt.Println("==========================")
}
