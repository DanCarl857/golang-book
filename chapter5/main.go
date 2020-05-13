package main

import "fmt"

func main() {
	x := [5]float64{98, 93, 77, 82, 83}

	total := 0.0
	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

	// Slices
	p := make([]float64, 5) // Create a slice associated with an array of length 5

	y := make([]float64, 5, 10) // Slice points to array of length 5, though underlying capacity is 10

	fmt.Println(p, y)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, slice2)

	slice3 := make([]int, 2)
	copy(slice3, slice2)

	fmt.Println(slice2, slice3)

	// Maps
	z := make(map[int]int)
	z[1] = 10
	fmt.Println(z[1])
}
