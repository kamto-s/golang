package main

import "fmt"

// const NamaLengkap = "Sukamto"

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(NamaLengkap)

	// string
	var youtube string = "Semangat Belajar Golang"
	fmt.Println(youtube)
	fmt.Printf("Tipe datanya adalah: %T \n", youtube)

	// boolean
	fmt.Println("=======================================")
	var subcribe bool = true
	fmt.Println(subcribe)
	fmt.Printf("Tipe datanya adalah: %T \n", subcribe)

	//uint8 uint16 uint32 uint64
	fmt.Println("=======================================")
	var angkaKecil uint8 = 255
	fmt.Println(angkaKecil)
	fmt.Printf("Tipe datanya adalah: %T \n", angkaKecil)

	//float32 float64
	fmt.Println("=======================================")
	var decimalKecil float32 = 255.1234567890
	fmt.Println(decimalKecil)
	fmt.Printf("Tipe datanya adalah: %T \n", decimalKecil)

	//int
	fmt.Println("=======================================")
	var intCoba int
	fmt.Println(intCoba)
	fmt.Printf("Tipe datanya adalah: %T \n", intCoba)

	//sting
	fmt.Println("=======================================")
	var stringCoba string
	fmt.Println(stringCoba)
	fmt.Printf("Tipe datanya adalah: %T \n", stringCoba)

	// implicite tipe/ tanpa deklarasi tipe data
	fmt.Println("=======================================")
	var impliciteTipe = 123
	fmt.Println(impliciteTipe)
	fmt.Printf("Tipe datanya adalah: %T \n", impliciteTipe)

	// multiple variable
	fmt.Println("=======================================")
	iniTipeApa := 345
	fmt.Println(iniTipeApa)
	fmt.Printf("Tipe datanya adalah: %T \n", iniTipeApa)

	// tipe data
	cekIni := 2.123456789
	fmt.Printf("Tipe data: %T \n", cekIni)

	fmt.Printf("Value: %v \n", cekIni)

	fmt.Printf("Tipe data: %T dan Valuenya: %v", cekIni, cekIni)

	var myCekIni string = fmt.Sprintf("Tipe data: %T dan Value: %v \n", cekIni, cekIni)

	fmt.Println(myCekIni)

	// ANGKA ===========================================

	// menampilkan true atau false
	fmt.Println("=======================================")
	fmt.Printf("True atau False: %t \n", 4 > 5)

	// menampilkan binary (base 2) dari sebuah angka
	fmt.Println("=======================================")
	fmt.Printf("Binary: %b \n", 10)

	// menampilkan octal (base 4) dari sebuah angka
	fmt.Println("=======================================")
	fmt.Printf("Octal: %o \n", 10)

	// menampilkan decimal (base 10) dari sebuah angka
	fmt.Println("=======================================")
	fmt.Printf("Decimal: %d \n", 10)

	// menampilkan hexadecimal (base 16) dari sebuah angka
	fmt.Println("=======================================")
	fmt.Printf("Hexadecimal: %x \n", 10)

	// notasi ilmiah
	fmt.Println("=======================================")
	fmt.Printf("Notasi Ilmiah: %e \n", 10.123456789)

	// menampilkan angka decimal dan dibulatkan
	fmt.Println("=======================================")
	fmt.Printf("Decimal bulatkan: %f \n", 10.129)

	// menampilkan decimal secara menyeluruh
	fmt.Println("=======================================")
	fmt.Printf("Decimal menyeluruh: %g \n", 10.123456789)

	// STRING ===========================================

	// string default
	fmt.Println("=======================================")
	fmt.Printf("Bahasa: %s \n", "Go")

	// string tanda kutip
	fmt.Println("=======================================")
	fmt.Printf("Bahasa: %q \n", "Golang")

	//  width dan precision
	fmt.Println("=======================================")
	fmt.Printf("%f \n", 100.1234)

	// menambahkan panjang ke kiri
	fmt.Println("=======================================")
	fmt.Printf("Bahasa: %40d \n", 99)

	// menambahkan panjang ke kanan
	fmt.Println("=======================================")
	fmt.Printf("Bahasa: %-40d \n", 99)

	// mengambil 2 digit dibelakang koma
	fmt.Println("=======================================")
	fmt.Printf("Ambil 2 digit dibelakang koma: %.2f \n", 100.123456789)

	// mengambil 2 digit dibelakang koma  plus tambah spasi
	fmt.Println("=======================================")
	fmt.Printf("Ambil 2 digit dibelakang koma plus tambah spasi: %20.2f cek \n", 100.123456789)

	// padding menambhakan karakter panjang
	fmt.Println("=======================================")
	fmt.Printf("No. Urut: %03d \n", 10)

	fmt.Println("=======================================")
	var noUrut string = fmt.Sprintf("%03d", 10)
	fmt.Println(noUrut)

	//
	fmt.Println("=======================================")
}

// Comma Ok Syntax ==========================
