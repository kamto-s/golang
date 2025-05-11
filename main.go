package main

import "fmt"

func main() {
	fmt.Println("==========================")
	fmt.Println("***BREAK")
	i := 1
	for i <= 10 {
		if i == 3 {
			fmt.Printf("---> Nilai ke-%v \n", i)
			break
		}

		fmt.Printf("Nilai ke-%v \n", i)
		i++
	}
	fmt.Println("baris ini tereksekusi setelah break")

	fmt.Println("==========================")

	// continue melanjutkan perulangan setelah kondisi tertentu
	fmt.Println("***CONTINUE")
	x := 1

	for x <= 10 {
		if x == 3 {
			fmt.Printf("---> Nilai ke-%v stop, tetapi tetap dilanjut dengan kondisi x += 7 \n", x)
			x += 7
			continue
		}

		fmt.Printf("Nilai ke-%v \n", x)
		x++
	}

	fmt.Println("==========================")
	fmt.Println("***GOTO")
	// goto adalah pernyataan yang digunakan untuk melompat ke label tertentu dalam kode
	y := 1

	for y <= 10 {

		if y == 5 {
			fmt.Printf("---> Nilai ke-%v stop, langsung loncat_kesini \n", y)
			goto loncat_kesini
		}

		fmt.Printf("Nilai ke-%v \n", y)
		y++
	}

loncat_kesini:
	fmt.Println("ini adalah label loncat_kesini")
}
