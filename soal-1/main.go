package main

import (
	"fmt"
	"sort"
)

func main() {
	Produk := []struct {
		Nama   string
		Harga  int
		Rating float32
		Likes  int
	}{
		{
			Nama:   "Indomie",
			Harga:  3000,
			Rating: 5,
			Likes:  150,
		},
		{
			Nama:   "Laptop",
			Harga:  4000000,
			Rating: 4.5,
			Likes:  123,
		},
		{
			Nama:   "Aqua",
			Harga:  3000,
			Rating: 4,
			Likes:  250,
		},
		{
			Nama:   "Smart TV",
			Harga:  4000000,
			Rating: 4.5,
			Likes:  42,
		},
		{
			Nama:   "Headphone",
			Harga:  4000000,
			Rating: 3.5,
			Likes:  90,
		},
		{
			Nama:   "Very Smart TV",
			Harga:  4000000,
			Rating: 3.5,
			Likes:  87,
		},
	}

	sort.SliceStable(Produk, func(i, j int) bool {
		return Produk[i].Harga < Produk[j].Harga
	})
	fmt.Println("Urutan Produk: ")
	fmt.Println(Produk)
}
