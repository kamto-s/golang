package main

import "fmt"

func main() {
	//operator relasional untuk membandingkan dua nilai
	// operator ini menghasilkan nilai boolean true atau false

	/*
		==  sama dengan
		!=  tidak sama dengan
		>   lebih besar dari
		<   lebih kecil dari
		>=  lebih besar dari sama dengan
		<=  lebih kecil dari sama dengan
	*/

	var angka1 int = 10
	var angka2 int = 20

	fmt.Printf("%d == %d: %t \n", angka1, angka2, angka1 == angka2)
	fmt.Printf("%d != %d: %t \n", angka1, angka2, angka1 != angka2)
	fmt.Printf("%d > %d: %t \n", angka1, angka2, angka1 > angka2)
	fmt.Printf("%d < %d: %t \n", angka1, angka2, angka1 < angka2)
	fmt.Printf("%d >= %d: %t \n", angka1, angka2, angka1 >= angka2)
}
