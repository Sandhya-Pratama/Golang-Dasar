package main

import "fmt"

func logging(){
	fmt.Println("Selesai melakukan fungsi")
}

func RunApp(){
	defer logging()
	fmt.Println("Aplikasi Selesai")
}

func main(){
	RunApp()
}
