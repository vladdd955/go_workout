package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	rrr := randomNumber(1, 150)
	fmt.Println("Make random number ", rrr[1], " ", rrr[2])

	res, err := count(rrr[1], rrr[2])
	if err != nil {
		log.Fatal("fatal bro.... error")
	}
	fmt.Println("And after plus first and second num, result is: ", res)
}

func count(x int, y int) (int, error) {
	var res int
	var err error
	res = x + y
	return res, err
}

func randomNumber(min int, max int) map[int]int {
	myMap := make(map[int]int)

	for i := 1; i <= 2; i++ {
		myMap[i] = mRand(min, max)
	}
	fmt.Println(myMap)
	return myMap
}

func mRand(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
