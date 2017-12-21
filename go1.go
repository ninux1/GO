package main

import "fmt"

func main() {
	list := []string{"Hello", "How", "Are", "You"}
	list = append(list, "Folks")
	list = append(list, "Doing")
	list = append(list, addString())
	for index, element := range list {
		fmt.Println(index, element)
	}
}

func addString() string {
	return "in this area"
}
