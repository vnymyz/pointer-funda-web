package main

import "fmt"

type Mahasiswa struct {
	Nama, Kelas, Npm string
}

func main() {
	data1 := Mahasiswa{"Vanya Mayazura", "4KA21", "16119468"} // tulis dengan npm kalian masing2
	data2 := data1

	data2.Nama = "Kobo Kanaeru" // kalau ini bebas isi apa

	fmt.Println(data1)
	fmt.Println(data2)
}
