package main

import (
	"fmt"
	"math/rand"
	"time"
)

type list []string

func main() {
	l := list{"One", "two", "three", "four", "five"}
	new := l.shuffle()
	for _, e := range new {
		fmt.Println(e)
	}
}

func (r list) shuffle() list {
	for index := range r {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(5)
		//swap(r, random, index)  golang does not require manual swapping steps.
		//Use the below line to swap in golang, Pretty quick.
		r[random], r[index] = r[index], r[random]
	}
	return r
}

/*
func swap(l list, rnd int, e int) {
	temp := l[rnd]
	l[rnd] = l[e]
	l[e] = temp
}
*/
