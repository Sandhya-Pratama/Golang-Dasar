package main

import "fmt"

func NamaKomplit() (awal, tengah, akhir string) {
	awal = "Sandhya"
	tengah = "Pratama"
	akhir ="Hutagalung"
	return awal, tengah, akhir
}

func main() {
	a, b, c := NamaKomplit()
	fmt.Println(a, b, c)
}