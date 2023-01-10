package main

import "fmt"

type UpdateMahasiswa struct {
	// buat field pada struct dengan Nama, Kelas, Npm bertipe String
}

func main() {
	var data1 UpdateMahasiswa = UpdateMahasiswa{"Vanya Mayazura", "4KA21", "16119468"} // isi data kalian
	var data2 *UpdateMahasiswa = &data1
	//tambahkan variabel dengan pointer

	data2 = &UpdateMahasiswa{".....", "....", "...."} // isi nilai yang kosong pada titik-titik bebas isi apa

	// cetak semua variabel data
	fmt.Println(...) 
	fmt.Println(...)
	fmt.Println(...)

	// buat lah variabel new dengan pointer dan parameter UpdateMahasiswa
	var data4 *UpdateMahasiswa = ...(...)
	data4.Npm = "...."
	fmt.Println(...)
}
