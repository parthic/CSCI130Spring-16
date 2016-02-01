package main

import "fmt"

func max(args ...int) int {
	var max int
	for _, a := range args {
		if a > max {
			max = a
		}
	}
	return max
}
func main() {
	fmt.Println(max(200, 85, 90, 576, 7896, 0000, 4738))
}
