package main

import "fmt"

func getName() (string, string) {
	return "sandhya", "pratama"
}

func main() {
	// awal, akhir := getName()
	// fmt.Println(awal, akhir)

	awal, _ := getName()
	fmt.Println(awal)
}