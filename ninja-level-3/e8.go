package main

import "fmt"

func e8() {
	switch {
	case false:
		fmt.Println("Should not print")
	case true:
		fmt.Println("Should print since switch doesn't have any expression")
	}
}
