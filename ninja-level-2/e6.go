package main

import "fmt"

//iota auto increments by 1

const (
	a4 = 2020 + iota
	b4 = 2020 + iota
	c4 = 2020 + iota
)

func main() {
	fmt.Println(a4) //Prints 2020
	fmt.Println(b4) //Prints 2021
	fmt.Println(c4) //Prints 2022
}
