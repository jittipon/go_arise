package idvalidator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func ValidateThaiId(s string) error {

	var result []string
	for _, char := range s {
		result = append(result, string(char))
	}
	fmt.Println("r", result)
	var sum = 0

	for i := 0; i < len(result)-1; i++ {
		number, _ := strconv.Atoi(result[i])
		sum = sum + number*(13-i)
		fmt.Printf("result[i] : %v || i: %v \n", result[i], 13-i)
	}

	var modResult = math.Mod(float64(sum), 11)

	var cal = 11 - modResult

	fmt.Println("sum", sum)
	fmt.Println("sum % 11: ", modResult)

	if len(result) < 13 {
		fmt.Println("id digits incorrect")
		return errors.New("id digits incorrect")
	}

	var lastIndex = result[len(result)-1]

	strValue := strconv.FormatFloat(cal, 'f', 0, 64)

	if strValue != lastIndex {
		fmt.Println("id incorrect")

		return errors.New("id incorrect")
	}

	return nil
}
