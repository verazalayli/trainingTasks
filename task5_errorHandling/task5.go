package task5_errorHandling

import (
	"errors"
	"fmt"
)

func Task5() {
	result, err := divide(3, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

func divide(a, b float64) (float64, error) {
	errZeroDivisible := errors.New("cannot divide zero")
	errZeroDivider := errors.New("cannot divide by zero")
	if a == 0 {
		return 0, errZeroDivisible
	}
	if b == 0 {
		return 0, errZeroDivider
	}
	result := a / b
	return result, nil
}
