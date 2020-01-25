package main

import "fmt"

func checkError(err error) bool {
	if err != nil {
		fmt.Println("Error occurred: ", err)
		return true
	}
	return false
}
