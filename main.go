package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv
	fmt.Println("===================================")
	fmt.Println("Konversi string ke int")
	var stringAngka string = "123"
	angkaInt, _ := strconv.ParseInt(stringAngka, 10, 64)
	fmt.Printf("Tipe data: %T \nValue: %v \n", angkaInt, angkaInt)

	fmt.Println("===================================")
	fmt.Println("Konversi string ke float")
	var stringAngka2 string = "10.123"
	angkaInt2, _ := strconv.ParseFloat(stringAngka2, 64)
	fmt.Printf("Tipe data: %T \nValue: %v \n", angkaInt2, angkaInt2)

	fmt.Println("===================================")
	fmt.Println("Konversi string ke bool")
	var stringAngka3 string = "0"
	angkaInt3, _ := strconv.ParseBool(stringAngka3)
	fmt.Printf("Tipe data: %T \nValue: %v \n", angkaInt3, angkaInt3)

	//
	//
	//
	fmt.Println("===================================")
}
