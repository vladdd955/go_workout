package main

import (
	"errors"
	"fmt"
	"log"
)

const notNecessary = "not_necessary"
const mustHave = "must_have"
const invalidCase = "invalid_case"

func main() {
	fmt.Println("heyyyy")

	myp := getMapForTest()
	result, err := preparedData(myp)
	if err != nil {
		log.Fatal("Prepare data are error")
	}

	fmt.Println("Final statement is : ", result)
}

func preparedData(myMap map[string][]string) (bool, error) {
	if myMap == nil {
		fmt.Println("myMap is nil")
		return false, errors.New("FATAL ERROR case empty map")
	}

	res := make(map[string]string)
	var count int
	for k, v := range myMap {
		for _, val := range v {
			fmt.Println(k)
			fmt.Println(val)
			if k == mustHave && val != "" {
				res[k] = val
			} else if k == notNecessary && val != "" {
				res[k] = val
			} else if k == invalidCase && val != "" {
				count++
				res[invalidCase] = val
			}
		}
	}

	if count != 0 {
		fmt.Println("How many invalid cases have inside map ", count)
	} else {
		fmt.Println("Do not have any invalid cases!")
	}

	var result bool
	if _, exists := res[invalidCase]; exists {
		result = false
	} else {
		result = true
	}

	return result, nil
}

func getMapForTest() map[string][]string {
	return map[string][]string{
		"must_have":     {"Take the bacon", "Take a something else"},
		"not_necessary": {"Fire", "lighter"},
		"invalid_case":  {"test case"},
	}
}
