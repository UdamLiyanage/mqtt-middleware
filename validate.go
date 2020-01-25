package main

import (
	"context"
	"fmt"
)

func validateRequest() {
	fmt.Println("Validate Request")
	var doc map[string]interface{}
	var res map[string]string
	response, err := database.Query(context.TODO(), "FOR d IN devices FILTER d.serial=='test_serial' RETURN d", doc)
	if err != nil {
		panic(err)
	}
	_, err = response.ReadDocument(context.TODO(), &res)
	if err != nil {
		println("Unidentified Device")
	}
	println("Device Identified! ", res["name"])
}
