package main

import (
	"fmt"
)

const CHARACTER string = "AAABBDDCEFGGG"

func main() {
	check, err := groupingCharacter(CHARACTER)
	if err != nil {
		fmt.Println("Error, something went wrong")
	}

	fmt.Println("Unique value from grouping function : ", check)
}

func groupingCharacter(char string) (string, error) {
	res := ""

	for i := 0; i < len(char); i++ {
		if i == 0 || char[i] != char[i-1] {
			res += string(char[i])
		}
	}

	return res, nil
}
