package main

import (
	"fmt"
)

//In Go for slices you do not mention the size while declaring it but not for slice
func e4() {
	var x = []int{1, 2, 4, 6, 8, 10, 12, 14, 16, 18} //Not mentioned the size here
	x = append(x, 53)
	fmt.Println(x)
	x = append(x, 54, 55, 56)
	fmt.Println(x)
	y := []int{57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}
