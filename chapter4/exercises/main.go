package main

import "fmt"

func main() {
	// Multiples of 3
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	// Fizz buzz
	for index := 1; index < 100; index++ {
		if index%3 == 0 && index%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if index%5 == 0 {
			fmt.Println("Buzz")
		} else if index%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(index)
		}
	}
}
