package main

import (
	"fmt"
)

func e4() {
	a2 := 42
	fmt.Printf("%d\t%b\t%#x", a2, a2, a2)
	b2 := a2 << 1 //Left shift
	fmt.Printf("%d\t%b\t%#x", b2, b2, b2)
}
