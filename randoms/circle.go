package main

import (
	"fmt"
	"math"
)

func main() {
	var area, radius float32
	const pi=3.14159 // use  math package to get math.Pi

	fmt.Println("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)

	area = pi * radius * radius
	fmt.Printf("The area of the circle with radius %.2f is %.2f\n", radius, area)

	area = math.Pi * radius * radius
	fmt.Printf("The area of the circle with radius %f is %f\n", radius, area)
	fmt.Printf("The area of the circle with radius %g is %g\n", radius, area)
}