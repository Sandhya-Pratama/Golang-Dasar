package main

import "fmt"

func main() {
	orang := map[string]string{
		"nama": "sandhya",
		"asal": "semarang",
		"pacar": "vivi",
	}

	fmt.Println(orang)
	fmt.Println(orang["nama"])
	fmt.Println(orang["asal"])

	buku := map[string]string{
		"penulis" : "sandhya ganteng",
		"judul" : "hujan dikala sendu",
		"tahun" : "2021",
	}
	fmt.Println(buku)

	//delete
	delete(buku, "tahun")
	fmt.Println(buku)
}