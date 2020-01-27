package main

import (
	"encoding/json"
	"testing"
)

func testValidateRequest(t *testing.T) {
	msg := map[string]string{
		"serial": "test_serial",
	}
	body, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	if !validateRequest(body) {
		t.Error("Device Serial: test_serial should be available")
	}
}

func testInvalidValidateRequest(t *testing.T) {
	msg := map[string]string{
		"serial": "test_",
	}
	body, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	if validateRequest(body) {
		t.Error("Device Serial: test_serial should not be available")
	}
}
