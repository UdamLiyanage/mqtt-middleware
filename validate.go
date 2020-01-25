package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
)

func validateRequest(message []byte) {
	var msg map[string]string
	fmt.Println("Validate Request")
	err := json.Unmarshal(message, &msg)
	if err != nil {
		panic(err)
	}
	if !validateSerial(msg["serial"]) {
		fmt.Println("Unidentified Device!")
		return
	}
	println("Device Identified! ")
}

func validateSerial(serial string) bool {
	var res map[string]string
	bindVars := map[string]interface{}{
		"serial": serial,
	}
	response, err := database.Query(context.TODO(), "FOR d IN devices FILTER d.serial==@serial RETURN d", bindVars)
	if err != nil {
		panic(err)
	}
	_, err = response.ReadDocument(context.TODO(), &res)
	if err != nil || driver.IsNoMoreDocuments(err) {
		return false
	}
	return true
}
