package main

import (
	"fmt"
	"strconv"
)

func Penjumlahan(bilanganPertama, bilanganKedua string) int {
	fmt.Print("Masukkan Besaran Bilangan Pertama: ")
	fmt.Scan(&bilanganPertama)
	fmt.Print("Masukkan Besaran Bilangan Kedua: ")
	fmt.Scan(&bilanganKedua)
	n1 := bilanganPertama
	n2 := bilanganKedua
	convertBilanganPertama, _ := strconv.Atoi(n1)
	convertBilanganKedua, _ := strconv.Atoi(n2)
	hasil := convertBilanganPertama + convertBilanganKedua
	fmt.Print("Hasil Penjumlahan Kedua Bilangan: ")
	return hasil
}

func Pengurangan(bilanganPertama, bilanganKedua string) int {
	fmt.Print("Masukkan Besaran Bilangan Pertama: ")
	fmt.Scan(&bilanganPertama)
	fmt.Print("Masukkan Besaran Bilangan Kedua: ")
	fmt.Scan(&bilanganKedua)
	n1 := bilanganPertama
	n2 := bilanganKedua
	convertBilanganPertama, _ := strconv.Atoi(n1)
	convertBilanganKedua, _ := strconv.Atoi(n2)
	hasil := convertBilanganPertama - convertBilanganKedua
	fmt.Print("Hasil Pengurangan Kedua Bilangan: ")
	return hasil
}

func Perkalian(bilanganPertama, bilanganKedua string) int {
	fmt.Print("Masukkan Besaran Bilangan Pertama: ")
	fmt.Scan(&bilanganPertama)
	fmt.Print("Masukkan Besaran Bilangan Kedua: ")
	fmt.Scan(&bilanganKedua)
	n1 := bilanganPertama
	n2 := bilanganKedua
	convertBilanganPertama, _ := strconv.Atoi(n1)
	convertBilanganKedua, _ := strconv.Atoi(n2)
	hasil := convertBilanganPertama * convertBilanganKedua
	fmt.Print("Hasil Perkalian Kedua Bilangan: ")
	return hasil
}

func Pembagian(bilanganPertama, bilanganKedua string) float64 {
	fmt.Print("Masukkan Besaran Bilangan Pertama: ")
	fmt.Scan(&bilanganPertama)
	fmt.Print("Masukkan Besaran Bilangan Kedua: ")
	fmt.Scan(&bilanganKedua)
	n1 := bilanganPertama
	n2 := bilanganKedua
	convertBilanganPertama, _ := strconv.ParseFloat(n1, 8)
	convertBilanganKedua, _ := strconv.ParseFloat(n2, 8)
	hasil := convertBilanganPertama / convertBilanganKedua
	fmt.Printf("Hasil Pembagian Kedua Bilangan: ")
	return hasil
}

func main() {
	fmt.Print("Pilih Operator Penjumlahan dibawah ini\n1. +\n2. - \n3. *\n4. /\n")
	var n int
	fmt.Print("Masukkan Pilihan Anda: ")
	fmt.Scan(&n)
	if n == 1 {
		fmt.Print("Anda Memilih Operator Penjumlahan\n")
		penjumlahan := Penjumlahan("", "")
		fmt.Print(penjumlahan)
	} else if n == 2 {
		fmt.Print("Anda Memilih Operator Pengurangan\n")
		Pengurangan := Pengurangan("", "")
		fmt.Print(Pengurangan)
	} else if n == 3 {
		fmt.Print("Anda Memilih Operator Perkalian\n")
		Perkalian := Perkalian("", "")
		fmt.Print(Perkalian)
	} else if n == 4 {
		fmt.Print("Anda Memilih Operator Pembagian\n")
		Pembagian := Pembagian("", "")
		fmt.Print(Pembagian)
	} else {
		fmt.Println("Masukkan Pilihan Anda Dengan Benar!!")
	}
}

// func main() {
// 	fmt.Print("Pilih Operator Penjumlahan dibawah ini\n1. +\n2. - \n3. *\n4. /\n")
// 	var n int
// 	fmt.Print("Masukkan Pilihan Anda: ")
// 	fmt.Scan(&n)
// 	if n == 1 {
// 		fmt.Print(penjumlahan)
// 		// var bilanganPertama int64
// 		// var bilanganKedua int64
// 		// fmt.Print("Masukkan Besaran Bilangan Pertama: ")
// 		// fmt.Scan(&bilanganPertama)
// 		// fmt.Print("Masukkan Besaran Bilangan Kedua: ")
// 		// fmt.Scan(&bilanganKedua)
// 		// hasil := bilanganPertama + bilanganKedua
// 		// fmt.Printf("Hasil Penjumlahan Kedua Bilangan Adalah = %d\n", hasil)
// 	} else if n == 2 {
// 		fmt.Print("Anda Memilih Operator Pengurangan\n")
// 		// var bilanganPertama int64
// 		// var bilanganKedua int64
// 		// fmt.Print("Masukkan Besaran Bilangan Pertama: ")
// 		// fmt.Scan(&bilanganPertama)
// 		// fmt.Print("Masukkan Besaran Bilangan Kedua: ")
// 		// fmt.Scan(&bilanganKedua)
// 		// hasil := bilanganPertama - bilanganKedua
// 		// fmt.Printf("Hasil Pengurangan Kedua Bilangan Adalah = %d\n", hasil)
// 	} else if n == 3 {
// 		fmt.Print("Anda Memilih Operator Perkalian\n")
// 		var bilanganPertama int64
// 		var bilanganKedua int64
// 		fmt.Print("Masukkan Besaran Bilangan Pertama: ")
// 		fmt.Scan(&bilanganPertama)
// 		fmt.Print("Masukkan Besaran Bilangan Kedua: ")
// 		fmt.Scan(&bilanganKedua)
// 		hasil := bilanganPertama * bilanganKedua
// 		fmt.Printf("Hasil Perkalian Kedua Bilangan Adalah = %d\n", hasil)
// 	} else if n == 4 {
// 		fmt.Print("Anda Memilih Operator Pembagian\n")
// 		var bilanganPertama float64
// 		var bilanganKedua float64
// 		fmt.Print("Masukkan Besaran Bilangan Pertama: ")
// 		fmt.Scan(&bilanganPertama)
// 		fmt.Print("Masukkan Besaran Bilangan Kedua: ")
// 		fmt.Scan(&bilanganKedua)
// 		hasil := bilanganPertama / bilanganKedua
// 		fmt.Printf("Hasil Pembagian Kedua Bilangan Adalah = %.2f\n", hasil)
// 	} else {
// 		fmt.Println("Masukkan Pilihan Anda Dengan Benar!!")
// 	}
// }
