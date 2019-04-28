package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for _, arg := range os.Args[1:] {
		s += " " + arg
	}
	fmt.Println(s)
}

/*
1.变量如果不需要使用，那么必须把它放到"_"中
*/
