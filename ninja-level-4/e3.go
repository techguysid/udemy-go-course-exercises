package main

import (
	"fmt"
)

//In Go for slices you do not mention the size while declaring it but not for slice
func e3() {
	var x = []int{1, 2, 4, 6, 8, 10, 12, 14, 16, 18} //Not mentioned the size here
	fmt.Println(x)
	fmt.Println(x[:5])  //Slicing the slice from 0 to index 5 .. index 5 not included
	fmt.Println(x[5:])  //Slicing from index 5 till end of slice
	fmt.Println(x[2:5]) //Slicing from index 2 (included) till index 5 (not included)
}
