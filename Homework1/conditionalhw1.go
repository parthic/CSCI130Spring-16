package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

/*type Person struct {
var Username string = "abc"
var Pwd string = "123"
}*/
const cond bool = true

func main() {
	var Username string = "abc"
	var Pwd string = "123"
	fmt.Printf("", Username+Pwd)

	t, exec := template.ParseFiles("c1.gohtml")
	logError(exec)
	exec = t.Execute(os.Stdout, cond)
	logError(exec)
}
func logError(exec error) {
	if exec != nil {
		log.Fatal(exec)
	}
}
