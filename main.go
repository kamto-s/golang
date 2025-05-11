package main

import "fmt"

func main() {
	angka := 4

	switch {
	case angka > 1:
		{
			fmt.Println("Angka lebih besar daripada 1")
			break
		}
	case angka > 2:
		{
			fmt.Println("Angka lebih besar daripada 2")
			break
		}
	default:
		{
			fmt.Println("Angka tidak dikenali")
		}
	}

	nilai := 40

	switch {
	case nilai >= 90 && nilai <= 100:
		{
			fmt.Println("Nilai A")
			break
		}
	case nilai >= 80 && nilai < 90:
		{
			fmt.Println("Nilai B")
			break
		}
	case nilai >= 70 && nilai < 80:
		{
			fmt.Println("Nilai C")
			break
		}
	default:
		{
			fmt.Println("Nilai Anda tidak memenuhi syarat")
			break
		}

	}
}
