package main

import "fmt"

func e4() {
	bd := 1992
	for {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
