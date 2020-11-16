package utils

import (
	"fmt"
	"strconv"
)

func Validate(input string) ([]string, error) {
	inputString := make([]string, 4)

	if len(input) == 4 {
		for i := 0; i <= 3; i++ {
			inputInt, _ := strconv.Atoi(string(input[i]))
			if inputInt != 0 || string(input[i]) == "0" {
				inputString[i] = string(input[i])
			}
		}

		if inputString[0] == inputString[1] || inputString[0] == inputString[2] || inputString[0] == inputString[3] || inputString[1] == inputString[2] || inputString[1] == inputString[3] || inputString[2] == inputString[3] {
			return nil, fmt.Errorf("same digits are not allowed")
		}

		return inputString, nil
	}

	return nil, fmt.Errorf("please enter 4 digits number")
}
