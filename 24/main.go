package main

import (
	"fmt"
	"level-1/24/point"
)

func main() {
	point1 := point.NewPoint(1, 1)
	point2 := point.NewPoint(2, 2)

	fmt.Println(point1.Distance(point2))
}
