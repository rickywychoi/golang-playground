package main

import "fmt"

func main() {
	type vroom int

	var a vroom = 1

	fmt.Println(a)
	fmt.Printf("%T", a)
}
