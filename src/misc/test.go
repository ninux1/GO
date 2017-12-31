package main

import "fmt"

func main() {
	c := color("Red")
	fmt.Println(c.describe("is an awesome color"))
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}

// This prints "Red is an awesome color".
// type can be declared anywhere in the golang program and not required to be declared always in the begining.
