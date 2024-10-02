package main

import "fmt"

func main() {
	// coba := 1
	// for coba <= 10{
	// fmt.Println("Perulangan ke", coba)
	// coba++}

	
	for coba := 1; coba <= 10; coba++{
	fmt.Println("Perulangan ke", coba)}

	fmt.Println("Selesai")

	names := []string{ "aki", "aku", "ako"}
	for i:=0; i< len(names); i++{
	fmt.Println(names[i])
	}

	for index, name := range names{
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names{
		fmt.Println(name)
	}
}