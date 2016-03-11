package main

import "fmt"

func main() {
	var num int
	fmt.Println("function takes an integer")
	fmt.Println("enter the value :  ")
	fmt.Scanf("%d", &num)
	//fmt.Println(val(num))
	val := func(n int) (rem int, even bool) {
		rem = n / 2
		if n%2 == 0 {
			//fmt.Printf("\n %d ,", rem)
			//fmt.Printf(" %t \n", true)
			even = true
		} else {
			//fmt.Printf("\n %d ,", rem)
			//fmt.Printf("%t \n", false)
			even = false
		}
		return rem, even
	}
	fmt.Println(val(num))
}
