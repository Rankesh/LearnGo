package main

import "fmt"

func main() {
	//Arrays
	var fruits [2]string
	//assign value
	fruits[0] = "Apple"
	fruits[1] = "orange"

	computers := [2]string{"Lenovo", "Apple"}
	companySlice := []string{"Infa", "AWS"}

	fmt.Println(computers)
	fmt.Println(len(companySlice))
}
