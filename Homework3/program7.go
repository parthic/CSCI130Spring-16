package main

import "fmt"

func main() {
	var n int
	var sum int = 0

	for n = 0; n < 1000; n++ {
		if n%3 == 0 {
			sum = sum + n
		} else if n%5 == 0 {
			sum = sum + n
		}
	}
	fmt.Println("the sum is", sum)
}
