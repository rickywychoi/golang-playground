package main

import (
	"fmt"
)

func main() {
	one()
	two()
	three()
	four()
	five()
}

func one() {
	fmt.Println("This is number one.")

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Println()
}

var x int
var y string
var z bool

func two() {
	fmt.Println("This is number two.")

	fmt.Println(x, y, z)
	fmt.Println(`These values are called "zero values".`)

	fmt.Println()
}

func three() {
	fmt.Println("This is number three.")

	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%d\t%s\t%t", x, y, z)
	fmt.Println(s)

	fmt.Println()
}

func four() {
	fmt.Println("This is number four.")

	type four int

	var x four

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	fmt.Println()
}

func five() {
	fmt.Println("This is number five.")

  type five int
  var y five = 5

	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", int(y), int(y))

	fmt.Println()
}
