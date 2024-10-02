package main

import "fmt"

func main() {
	var nilai_ujian = 90
	var remidi = 80

	var ujian_lulus bool = nilai_ujian > 80
	var remidi_lulus bool = remidi > 80

	var lulus bool = ujian_lulus && remidi_lulus

	fmt.Println(lulus)
}
