package main

import (
	"fmt"
)

//In Go for arrays you mention the size while declaring it but not for slice
func e1() {
	var x = [5]int{1, 2, 4, 6, 8}
	fmt.Println(x)
	for _, v := range x {
		fmt.Printf("%v\t%T\n", v, v)
	}
}
