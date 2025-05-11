package main

import "fmt"

func main() {
	/*
			+ penambahan
			- pengurangan
			* perkalian
			/ pembagian
			% modulus
		++ increment (menambah 1)
		-- decrement (mengurangi 1)
	*/

	var angka1 int = 10
	var angka2 int = 5

	penjumlahan := angka1 + angka2
	pengurangan := angka1 - angka2
	perkalian := angka1 * angka2
	pembagian := angka1 / angka2
	modulus := angka1 % angka2
	angka1++
	angka2--

	fmt.Println("Hasil penjumlahan: ", penjumlahan)
	fmt.Println("Hasil pengurangan: ", pengurangan)
	fmt.Println("Hasil perkalian: ", perkalian)
	fmt.Println("Hasil pembagian: ", pembagian)
	fmt.Println("Hasil modulus: ", modulus)
	fmt.Println("Hasil increment +1: ", angka1)
	fmt.Println("Hasil decrement -1: ", angka2)
}
