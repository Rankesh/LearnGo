package main

import "fmt"

func main() {

	//long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	//shot method
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	//FizzBuzz
	for x := 1; x <= 100; x++ {
		if x%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if x%3 == 0 {
			fmt.Println("Fizz")
		} else if x%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(x)
		}
	}
}
