package utils

import "fmt"

var version = "1.0"

func Print(s string) {
	fmt.Println(version + ":" + s)
}
