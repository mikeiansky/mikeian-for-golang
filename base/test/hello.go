package test

import (
	"errors"
	"fmt"
)

func hello(msg string) (string, error) {
	fmt.Println("hello:", msg)
	if 1 == 1 {
		return "", errors.New("test error")
	}
	return "hello " + msg, nil
}
