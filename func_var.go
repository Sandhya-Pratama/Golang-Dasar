package main

import "fmt"
func penambahantotal (number ...int)int{
	total := 0
	for _, number:= range number{
		total+=number
	}
	return total
}
func main() {
	total := penambahantotal(10, 10, 10, 10, 10)
	fmt.Println(total)
}