package main

import "fmt"

func main() {
	name := "jokonti"

	if name == "sandhya" {
		fmt.Println("benar ini sandhya")
	} else if name=="joko"{
		fmt.Println("benar ini joko")
	}else {
		fmt.Println("bukan sandhya")
	} 

	// panjang:=len(name)
	// if panjang>5{
	// 	fmt.Println("nama teralu panjang")
	// } else {
	// 	fmt.Println("nama cukup")
	// }

	if panjang:=len(name); panjang>5{
		fmt.Println("nama teralu panjang")
	} else {
		fmt.Println("nama cukup")
	}
}