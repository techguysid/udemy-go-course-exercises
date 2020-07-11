package main

import (
	"fmt"
)

//In Go for slices you do not mention the size while declaring it but not for slice
func e2() {
	var x = []int{1, 2, 4, 6, 8} //Not mentioned the size here
	fmt.Println(x)
	for _, v := range x {
		fmt.Printf("%v\t%T\n", v, v)
	}
}
