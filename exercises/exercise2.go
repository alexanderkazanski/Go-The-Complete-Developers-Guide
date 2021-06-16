package main

import (
	"os"
	"fmt"
)

func main() {
	file := os.Args[1]
	data, _ := os.Open(file)

	bs := make([]byte, 99999)
	data.Read(bs)
	fmt.Println(string(bs))
}