package main

import "fmt"

func main() {
	nama_orang := [...]string{"Sandhya", "Tama", "Huta", "Sendok", "Galung", "Stefanus", "Ayam", "yusup"}

	slice1 := nama_orang[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := nama_orang[:4]
	fmt.Println(slice2)


	slice3 := nama_orang[4:]
	fmt.Println(slice3)

	//manual ini
	var slice4 []string = nama_orang[:]
	fmt.Println(slice4)

	// dan ganti index slice
	day := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	dayslice1 := day[5:]
	fmt.Println(day)
	fmt.Println(dayslice1)

	dayslice1[0]="sabtu baru"
	dayslice1[1]="minggu baru"
	fmt.Println(day)
	fmt.Println(dayslice1)

	//coba untuk append
	dayslice2 := append(dayslice1, "libur baru")
	fmt.Println(dayslice2)

	//make slice
	newslice := make([]string, 2, 5)
	newslice[0]="budi"
	newslice[1]="yanto"

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	//COPY SLICE
	fromslice :=day[:]
	toslice :=make([]string, len(fromslice), cap(fromslice))

	copy(toslice, fromslice)

	fmt.Println(fromslice)
	fmt.Println(toslice)

}