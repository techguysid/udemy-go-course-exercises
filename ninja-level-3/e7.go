package main

import "fmt"

func e7() {
	x := "Siddharth"

	if x == "IDontKnow" {
		fmt.Println("Dont know what x is")
	} else if x == "Siddharth" {
		fmt.Println("Ok so x is Siddharth")
	} else {
		fmt.Println("Seriously don't know what x is")
	}
}
