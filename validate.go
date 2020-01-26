package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"net/http"
)

func validateRequest(message []byte) bool {
	var msg map[string]string
	err := json.Unmarshal(message, &msg)
	checkError(err)
	if !validateSerial(msg["serial"]) {
		fmt.Println("Unidentified Device!")
		return false
	}
	if err != nil {
		panic(err)
	}
	_, err = http.Post("http://127.0.0.1:9000/mqtt-broker", "application/json", bytes.NewBuffer(message))
	if err != nil {
		panic(err)
	}
	return true
}

func validateSerial(serial string) bool {
	var res map[string]string
	bindVars := map[string]interface{}{
		"serial": serial,
	}
	response, err := database.Query(context.TODO(), "FOR d IN devices FILTER d.serial==@serial RETURN d", bindVars)
	checkError(err)
	_, err = response.ReadDocument(context.TODO(), &res)
	if checkError(err) || driver.IsNoMoreDocuments(err) {
		return false
	}
	return true
}
