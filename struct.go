package main

import "fmt"

type Customer struct {
	Nama, Adress string
	Age          int
}

func main() {
	var eko Customer
	fmt.Println(eko)

	eko.Nama = "Eko puki"
	eko.Adress = "jepang"
	eko.Age = 19
	fmt.Println(eko)
}