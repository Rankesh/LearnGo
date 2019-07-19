package main

import "fmt"

func main() {

	//define map
	emails := make(map[int]string)

	emails[1] = "bob@gmail.com"
	emails[2] = "sharon@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails[2])

	//Delete

	delete(emails, 1)
	fmt.Println(emails)

	//Shorter Assignment
	companies := map[string]string{"RK": "RK@GK.com", "BK": "BK@GK.COM"}
	fmt.Println(companies)
}
