package main

import "fmt"

var a int = 42
var b bool = true
var c string = "James Bond"

func e3() {
	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)
}
