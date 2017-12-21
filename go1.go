package main

import "fmt"

func main() {
	list := []string{'Hello', 'How', 'Are', 'You'}
	list = append(list, 'Folks')
	list = append(list, 'Doing')
	for index, element := range list {
		fmt.Println(index, element)
	} 
}