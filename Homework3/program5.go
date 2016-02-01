package main

import "fmt"

func main() {
	var n int
	for n = 0; n <= 100; n++ {
		if n%2 == 0 {
			fmt.Println(n)
		} else {
			fmt.Printf(" ")
		}
	}
}
