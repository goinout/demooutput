package main

import (
	"fmt"

	"github.com/goinout/goinout"
)

var Handle goinout.OutputFunc

func init() {
	Handle = output
}

func output(params map[string]interface{}, data []byte) error {
	fmt.Println("params", params)
	fmt.Println("data", data)
	return nil
}
