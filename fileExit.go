package main

import (
	"os"
	"fmt"
)

func main(){

	var l bool
	l = FileExists(os.Args[1])

	fmt.Println(l)
}
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
