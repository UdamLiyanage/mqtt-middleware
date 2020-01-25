package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func validateRequest(message []byte) {
	var msg map[string]string
	fmt.Println("Validate Request")
	err := json.Unmarshal(message, &msg)
	if err != nil {
		panic(err)
	}
	var doc map[string]interface{}
	var res map[string]string
	response, err := database.Query(context.TODO(), "FOR d IN devices FILTER d.serial=='"+msg["serial"]+"' RETURN d", doc)
	if err != nil {
		panic(err)
	}
	_, err = response.ReadDocument(context.TODO(), &res)
	if err != nil {
		println("Unidentified Device")
	}
	println("Device Identified! ", res["name"])
}
