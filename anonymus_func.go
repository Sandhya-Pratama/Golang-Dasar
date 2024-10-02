package main

import "fmt"

type Blacklist func(string) bool

func RegisterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("kamu terblokir", name)
	} else{
		fmt.Println("selamat datang" ,name)
	}
}

func main(){
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	RegisterUser("yanto", blacklist)
}