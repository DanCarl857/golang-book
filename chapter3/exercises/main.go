package main

import "fmt"

func main() {
	var d float64
	var c float64
	fmt.Println("Enter a degree in fahrenheit: ")
	fmt.Scanf("%f", &d)

	c = (d - 32) * 5 / 9
	fmt.Println(c)

	var f float64
	fmt.Println("Now converting feet into metres,\nEnter value in feet: ")
	fmt.Scanf("%f", &f)

	m := f * 0.3048
	fmt.Println(m)
}
