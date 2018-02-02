package main

import "fmt"

type cylinder struct {
	len float64
	breadth	float64
	height	float64
}

func main() {
	cyl := cylinder{2.4, 3.5, 5.6}
	cyl1 := cylinder{8.4, 8.5, 6.6}

	//a := Area(cyl.len, cyl.breadth, cyl.height)
	//fmt.Println("Cylinder volume before modification is", a)

	//fmt.Println("changed value Area by reference ", cyl.change())

	//a = Area(cyl.len, cyl.breadth, cyl.height)
	//fmt.Println("Cylinder volume before modification is", a)

	fmt.Println("Area of Cyl is ", cyl.Area())
	fmt.Println("Area of cyl 1 is", cyl1.Area())
}


/*
func (c *cylinder)Area() float64{
	//c.len = 5.6
	//c.breadth = 6.6
	//c.height = 3.5
	return c.len * c.breadth * c.height
}
*/

func (c cylinder)Area() float64 {
	return c.len * c.breadth * c.height
}

/*
func Area(l float64, b float64, h float64) float64 {
	return l * b * h
}
*/