package main

import (
	"fmt"
)

// Handle the output
func Handle(params map[string]interface{}) error {
	fmt.Println("params", params)
	return nil
}
