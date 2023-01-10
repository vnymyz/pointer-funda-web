package main

import "fmt"

type Address struct {
	City, Province, Country string // ini skema atau model nya
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // address 2 pointer ke address 1
	// kita coba buat address 3
	var address3 *Address = &address1

	// kita coba lakukan perubahan variable atau assign ulang adddress 2
	*address2 = Address{"Jakarta", "Bogor", "Depok"} // address 2 akan tetep address 2 karena pointer nya

	// karena kita menggunakan * jadi nanti akan mengacu pada address 2 Address nya
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// membuat function new
	var adddress4 *Address = new(Address)
	adddress4.City = "Surabaya"
	fmt.Println(adddress4)
}

// soal 3 jelaskan logika dari program ini
