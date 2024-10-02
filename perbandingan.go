package main

import "fmt"

func main() {
	var name1 = "Sandhya"
	var name2 = "Sandhya"
	var a = 10
	var b = 2
	var c = 12


	var pertambahan bool= a+b == c
	var tau bool= a>b
	var result bool = name1 != name2

	fmt.Println(result)
	fmt.Println(pertambahan)
	fmt.Println(tau)
}
