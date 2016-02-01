package main

import "fmt"

func main() {
	var n int
	fmt.Println(0)
	fmt.Println(1)
	fmt.Println(2)

	for n = 0; n <= 100; n++ {
		if (n%3 == 0) && (n%5 == 0) {
			fmt.Println("FizzBuzz")
		} else if n%3 == 0 {
			fmt.Println("Fizz")
		} else if n%5 == 0 {
			fmt.Println("Buzz ")
		} else {
			fmt.Println(n)
		}
	}
}
