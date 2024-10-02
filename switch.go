package main

import "fmt"

func main() {
	name := "sandhya"

	switch name{
	case "sandhya":
		fmt.Println("benar ini sandhya")
	case "joko":
		fmt.Println("bukan ini sandhya")
	default:
		fmt.Println("lah gatau siapa")
	} 

	panjang := len(name)

	switch panjang>5 {
	case true:
		fmt.Println("benar")
	case false:
		fmt.Println("salah")	
	}
}