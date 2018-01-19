package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ffff0",
		"green": "#4bf90",
		"white": "#fffff",
	}
	changeMap(colors)
	printMap(colors)
	//fmt.Println(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

func changeMap(c map[string]string) {
	c["yellow"] = "#000dfde"
}
