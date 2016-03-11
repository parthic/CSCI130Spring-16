package main

import "fmt"

func main() {
	var a, b int
	fmt.Printf("enter value of largest number:  ")
	fmt.Scanf("%d\n", &a)
	fmt.Printf("enter value of smallest number: ")
	fmt.Scanf("%d", &b)
	fmt.Println("the remainder is", a%b)
}
