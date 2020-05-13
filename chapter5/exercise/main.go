package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	p := make([]int, 3, 9)
	fmt.Println(len(p))

	smallest := 100

	for _, value := range x {
		if value < smallest {
			smallest = value
		}
	}

	fmt.Println(smallest)
}
