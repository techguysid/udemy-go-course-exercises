package main

import (
	"fmt"
)

//In Go for slices you do not mention the size while declaring it but not for slice
func main() {
	var x = []int{1, 2, 4, 6, 8, 10, 12, 14, 16, 18} //Not mentioned the size here
	x = append(x[:4], x[6:]...)                      //Deleting from slice..using append and slicing the slice
	fmt.Println(x)
}
