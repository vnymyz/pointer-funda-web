package main

import "fmt"

type Address struct {
	City, Province, Country string // ini skema atau model nya
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"} // address 1 akan tetep address 1 outputnya
	address2 := &address1                                    // address 2 pointer ke address 1
	// address 2 pointer
	address2.City = "Bandung" // ini field

	// kita coba lakukan perubahan variable atau assign ulang adddress 2
	address2 = &Address{"Jakarta", "Bogor", "Depok"} // address 2 akan tetep address 2 karena pointer nya
	// jadi disini ceritanya nanti si address 2 akan membuat data baru
	// sehingga address 1 tidak ikut berubah

	fmt.Println(address1)
	fmt.Println(address2)
}

// soal nomor 2
// buatlah data struct dengan nama npm kelas lalu ubahlah data awal mahasiswa menjadi data yang terbaru
//menggunakan pointer
