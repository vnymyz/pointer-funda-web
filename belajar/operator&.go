package main

import "fmt"

type Address struct {
	City, Province, Country string // ini skema atau model nya
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"} // ini dia ngikutin struct nya
	address2 := &address1                                    // disini menunjukkan bahwa pointer nya address 1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2) // prinsip pointer dalam golang jika 1 nilai variable diubah maka yg kedua yg di refer akan berubah jg
}
