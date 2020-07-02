package main

import "fmt"

type hotdog1 int

var x2 hotdog1
var y2 int

func main() {
	fmt.Println(x2)
	fmt.Printf("%T\n", x2)
	x2 = 42
	fmt.Println(x2)
	y2 = int(x2)
	fmt.Println(y2)
	fmt.Printf("%T\n", y2)
}
