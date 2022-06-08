package models

type Produk struct {
	ID     string  `json:"id"`
	Nama   string  `json:"nama"`
	Detail string  `json:"detail"`
	Harga  int     `json:"harga"`
	Rating float32 `json:"rating"`
	Likes  int     `json:"likes"`
}
