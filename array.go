package main

import "fmt"

func main() {
	// coba type diubah
	type Nama string
	type Sayang int

	var name [4]Nama

	name[0] = "sandhya"
	name[1] = "pratama"
	name[2] = "hutagalung"
	name[3] = "cassano"

	fmt.Println(name[0], name[1], name[2], name[3])
/// coba integer
	var nilai = [4]Sayang{
		10,
		20,
		30,
		40,
	}

	fmt.Println(nilai)
	//panjang
	fmt.Println(len(nilai))

	//ubah isi index
	nilai[0] = 50
	fmt.Println(nilai[0])
}