package main

import "fmt"

func getGoobye(name string) string {
	return "hello, and goodbye " + name
}

func main() {
	bye := getGoobye
	fmt.Println(bye("Eko"))
}