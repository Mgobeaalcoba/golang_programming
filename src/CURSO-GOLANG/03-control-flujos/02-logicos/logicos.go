package main

import "fmt"

func main() {
	fmt.Println("Section NOT")
	// Not: Su operador es !
	fmt.Println(!true)
	// false

	fmt.Println("Section AND")
	// And: Su operador es &&
	//fmt.Println(true && true)
	// true
	fmt.Println(true && false)
	// false
	//fmt.Println(false && false)
	// false

	fmt.Println("Section OR")
	// Or: Su operador es ||
	//fmt.Println(true || true)
	// true
	fmt.Println(true || false)
	// true
	//fmt.Println(false || false)
	// false

	fmt.Println("Section TEST")
	b := 2

	r := b == 2 && (4 > b)

	fmt.Println(r)
	// true

	r = b == 2 && !(4 > b)

	fmt.Println(r)
	// false

	r = b == 2 || !(4 > b)

	fmt.Println(r)
	// true

}