package utils

import "fmt"

var version = "2.0"

func Print(s string) {
	fmt.Println(version + ":" + s)
}
