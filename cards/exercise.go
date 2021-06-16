package main

import "fmt"

func exercise() {
	var list = []int{0, 1, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, l := range list {
		if l%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
