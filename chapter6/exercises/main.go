package main

import "fmt"

func evenOrOddFraction(val int) (int, bool) {
	if val%2 == 0 {
		return val / 2, true
	}
	return val / 2, false
}

func largestNumber(args ...int) int {
	largest := 0
	for _, value := range args {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	fmt.Println("Even or Odd Fraction")
	fmt.Println(evenOrOddFraction(2))

	fmt.Println("Generator which generates Odd Numbers")
	nextOdd := makeOddGenerator()

	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	fmt.Println("Fibonacci for n = 10")
	fmt.Println(fibonacci(10))

	x := 1
	y := 2

	swap(&x, &y)

	fmt.Println("Swapping")
	fmt.Println(x, y)
}
