package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	var name = "Rankesh"
	var age int32 = 32
	const isCool = true
	var z = cmplx.Sqrt(-5+12i) 

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", age)
	fmt.Println(z)
}
