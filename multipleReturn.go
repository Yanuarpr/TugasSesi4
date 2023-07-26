package main

import (
	"fmt"
)

// Fungsi untuk menemukan nilai maksimum dan minimum
func findMaxAndMin(numbers ...int) (max, min int) {
	// Inisialisasi nilai maksimum dan minimum dengan angka pertama
	max = numbers[0]
	min = numbers[0]

	// Loop melalui angka-angka inputan
	for _, num := range numbers {
		// Temukan nilai maksimum
		if num > max {
			max = num
		}
		// Temukan nilai minimum
		if num < min {
			min = num
		}
	}

	return
}

func main() {
	// Inputkan 6 angka
	var a, b, c, d, e, f int
	fmt.Print("Masukkan 6 angka dipisahkan dengan spasi: ")
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	// Panggil fungsi untuk menemukan nilai maksimum dan minimum
	maximum, minimum := findMaxAndMin(a, b, c, d, e, f)

	// Dereferencing pointer (dalam contoh ini, tidak ada penggunaan pointer)
	fmt.Printf("Nilai maksimum: %d\n", maximum)
	fmt.Printf("Nilai minimum: %d\n", minimum)
}
