package main

import "fmt"

type UpdateMahasiswa struct {
	Nama, Kelas, Npm string
}

func main() {
	data1 := UpdateMahasiswa{"Vanya Mayazura", "4KA21", "16119468"} // tulis data kalian masing2
	data2 := &data1

	data2 = &UpdateMahasiswa{"Muhammad Ibnu", "2KA26", "999999"} // ini bebas isi apa

	fmt.Println(data1)
	fmt.Println(data2)
}
