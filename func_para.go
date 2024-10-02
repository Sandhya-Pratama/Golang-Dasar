package main

import "fmt"

func Fillterword(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func filterkasar(name string) string  {
	if name == "anjing"{
	return "..."
	} else{
	return name
	}
}

func main(){
	Fillterword("goblok", filterkasar)

	Filter:= filterkasar
	Fillterword("eko", Filter)
}