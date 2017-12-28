package main

import "fmt"

type myarr []int

func (a myarr) add() int {
	sum := 0
	for index, element := range a {
		fmt.Println(index)
		sum += element
	}
	return sum
}
