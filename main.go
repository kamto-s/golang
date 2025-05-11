package main

import (
	"errors"
	"fmt"
)

func penjumlahan(angka1, angka2 int) (int, error) {
	if angka1 > 10 || angka2 > 10 {
		return 0, errors.New("angka tidak boleh lebih dari 10")
	}

	return angka1 + angka2, nil
}

func main() {
	hasil, err := penjumlahan(9, 10)

	fmt.Println("Hasil penjumlahan:", hasil)
	fmt.Println(err)
}
