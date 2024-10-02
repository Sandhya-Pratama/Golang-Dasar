package main

import "fmt"

func getHello(nama string) string {
	hello := "hai " + nama
	return hello
}

func main() {
	hasil := getHello("eko")
	fmt.Println(hasil)
}