package main

import "fmt"

func main(){
	type NoHP string

	var  HPsandhya NoHP = "082123456"

	var contohHP = "08212121212"
	var NoHPBaru NoHP= NoHP(contohHP)

	fmt.Println(HPsandhya)
	fmt.Println(NoHPBaru)
}