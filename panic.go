package main

import "fmt"

func endapp(){
	fmt.Println("Aplikasi berhenti")
	massage:= recover()
		fmt.Println("terjadi error", massage)
}

func runApp(error bool){
	defer endapp()
	if error{
		panic("error cuk")}
}

func main(){
	runApp(true)
	fmt.Println("sandhya")
}
