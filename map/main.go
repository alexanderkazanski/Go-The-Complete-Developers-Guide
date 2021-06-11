package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	goThroughColors(colors)
}

func goThroughColors(c map[string]string) {
	for key, value := range c {
		fmt.Println(key)
		fmt.Println(value)
	}
}
