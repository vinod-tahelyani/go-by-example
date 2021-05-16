package main

import (
	"errors"
	"fmt"
	"os"
)

func Errors()  {

	err1 := getErrorMessage(2)

	fmt.Println(err1)

	_, err2 := f1(42)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(2)
	}



}

func f1(val int) (int, error) {
	if val == 42 {
		return -1, errors.New("can't work with 42")
	}
	return val+1, nil
}

func getErrorMessage(code int) error {
	switch code {
	case 1:
		return &argError{
			errorMessage: "abc",
			errorCode: 1,
		}
	case 2:
		return &argError{
			errorMessage: "def",
			errorCode: 2,
		}
	case 3:
		return &argError{
			errorMessage: "ghi",
			errorCode: 3,
		}
	default:
		return &argError{
			errorMessage: "abc",
			errorCode: -1,
		}
	}
}


type argError struct {
	errorMessage string
	errorCode int
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s - %d", e.errorMessage, e.errorCode)
}