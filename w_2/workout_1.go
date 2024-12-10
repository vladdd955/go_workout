package main

import "fmt"

var checkStatement bool
var p *int

func main() {
	tryPointer := 42
	ttt := 33
	fillStatement()

	if checkStatement {
		p = &tryPointer
	} else {
		p = &ttt
	}

	fmt.Println(*p)
}

func fillStatement() {
	checkStatement = false
}
