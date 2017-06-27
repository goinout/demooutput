package main

import (
	"fmt"
)

func output(params map[string]interface{}) error {
	fmt.Println("params", params)
	return nil
}
