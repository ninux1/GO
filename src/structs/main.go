package main

import "fmt"

type empdetails struct {
	empcity    string
	empcountry string
}

type employee struct {
	empid   int
	empname string
	empd    empdetails
}

func main() {
	o1 := employee{empid: 12345,
		empname: "Ninad",
		empd: empdetails{
			empcity:    "Portland",
			empcountry: "USA",
		},
	}

	fmt.Println("Before Update")
	o1.print()
	fmt.Println()
	o1.updatename()
	fmt.Println("After Update")
	o1.print()
}

func (e employee) print() {
	fmt.Printf("%+v", e)
}

func (e *employee) updatename() {
	(*e).empname = "Siddhi"
}
