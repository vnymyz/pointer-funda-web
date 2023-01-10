package main

import "fmt"

type Address struct {
	City, Province, Country string // ini skema atau model nya
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"} // ini dia ngikutin struct nya
	address2 := address1                                     // jadi address 1 datanya di clone atau duplikat ke address 2

	address2.City = "Bandung"

	fmt.Println(address1) // tidak berubah address 1 karena golang
	// semisal kita pake bahasa lain java atau php mustinya address 1 berubah
	fmt.Println(address2)
}

// soal nomor 1
// buatlah data struct dimana berisi nama npm dan kelas kalian dan menggunakan function untuk memanggil 2 variable
// dimana salah satu variable nya memiliki isi data yang berbeda pada model struct yang sudah ada
