package main

import "fmt"

func main() {
	//percabangan if else
	nilai := 65

	if nilai >= 90 && nilai <= 100 {
		fmt.Println("Nilai Anda A")
	} else if nilai >= 80 && nilai < 90 {
		fmt.Println("Nilai Anda AB")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("Nilai Anda B")
	} else if nilai >= 60 && nilai < 70 {
		fmt.Println("Nilai Anda C")
	} else {
		fmt.Println("Anda Tidak Lulus")
	}
}
