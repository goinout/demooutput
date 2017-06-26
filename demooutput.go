package main

import (
	"fmt"
)

// Handle the out plugin
func Handle(params map[string]interface{}, data []byte) error {
	fmt.Println("params", params)
	fmt.Println("data", data)
	return nil
}
