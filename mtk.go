package MTK

import "fmt"

func Tambah(a int, b int) int {
	return a + b
}

func Kurang(a int, b int) int {
	return a - b
}

func Kali(a int, b int) int {
	return a * b
}

func Bagi(a int, b int) int {
	if b == 0 {
		fmt.Println("Tidak bisa dibagi 0")
	}
	return a / b
}
