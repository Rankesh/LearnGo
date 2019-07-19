package main

import "fmt"

func main() {

	ids := []int{33, 76, 54, 23, 11, 2}

	//Loop through ids using range

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d\n", sum)

	//define map
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
